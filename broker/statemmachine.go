package key

import (
	"encoding/binary"
	"errors"
	"event-stream/protocol"
	"event-stream/util"
	sm "github.com/lni/dragonboat/v4/statemachine"
	"github.com/robfig/cron/v3"
	"github.com/tecbot/gorocksdb"
	"google.golang.org/protobuf/proto"
	"io"
	"os"
	"path/filepath"
	"time"
)

var processedIndex = []byte("processedIndex")

const current = "current"
const snapshot = "snapshot"
const lasProcessIndex = "lasProcessIndex"

type EventStateMachine struct {
	// 快照目录
	checkpointDir string
	shardID       uint64
	replicaID     uint64
	db            *gorocksdb.TransactionDB
	responses     chan sm.Result
	// 数据存储目录
	dir string
}

func NewEventStateMachine(dir string, shardID uint64, replicaID uint64) *EventStateMachine {
	return &EventStateMachine{dir: dir, shardID: shardID, replicaID: replicaID}
}

func (s *EventStateMachine) Open(stopc <-chan struct{}) (uint64, error) {
	db, err := createDb(s.dir)
	if err != nil {
		return 0, err
	}
	s.db = db
	slice, err := db.Get(gorocksdb.NewDefaultReadOptions(), []byte(lasProcessIndex))
	if err != nil || slice.Size() == 0 {
		return 0, err
	}
	return binary.LittleEndian.Uint64(slice.Data()), nil

}

func (s *EventStateMachine) Update(entries []sm.Entry) (res []sm.Entry, err error) {

	for i := range entries {
		entry := entries[i]
		process(s.db, &entry)
		s.responses <- entry.Result
	}

	return entries, nil
}

func process(db *gorocksdb.TransactionDB, entry *sm.Entry) error {
	var event protocol.Event
	if err := proto.Unmarshal(entry.Cmd, &event); err != nil {
		return err
	}

	switch event.Type {
	case protocol.EventType_JOB_CREATE:
		var job protocol.JobCreate
		if err := proto.Unmarshal(event.Value, &job); err != nil {
			return err
		}
		schedule, err := cron.ParseStandard(job.Cron)
		if err != nil {
			return err
		}
		slice, err := db.Get(gorocksdb.NewDefaultReadOptions(), []byte(job.Name))
		if err != nil {
			return err
		}
		if len(slice.Data()) > 0 {
			return errors.New("job存在")
		}
		next := schedule.Next(time.Now())
		err = db.Put(gorocksdb.NewDefaultWriteOptions(), []byte(job.Name), event.Value)
		if err != nil {
			return err
		}
		nanosecond := next.Nanosecond()
		err = db.Put(gorocksdb.NewDefaultWriteOptions(), Int64Byte(uint64(nanosecond)), event.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *EventStateMachine) Lookup(key interface{}) (interface{}, error) {
	slice, err := s.db.Get(gorocksdb.NewDefaultReadOptions(), key.([]byte))
	if err != nil {
		return nil, err
	}
	return slice.Data(), nil
}

func (s *EventStateMachine) Sync() error {
	return nil
}

func (s *EventStateMachine) PrepareSnapshot() (interface{}, error) {
	checkpoint, err := s.db.NewCheckpoint()
	if err != nil {
		return nil, err
	}
	if err = checkpoint.CreateCheckpoint(s.checkpointDir, 64); err != nil {
		return nil, err
	}
	return checkpoint, nil

}

func (s *EventStateMachine) SaveSnapshot(point interface{}, writer io.Writer, notify <-chan struct{}) error {
	checkpoint := point.(*gorocksdb.Checkpoint)
	defer checkpoint.Destroy()
	//不通过快照逐条去读取记录，直接拷贝快照文件
	bytes, err := util.Zip(s.checkpointDir)
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (s *EventStateMachine) RecoverFromSnapshot(reader io.Reader, i <-chan struct{}) error {

	s.db.Close()
	if err := util.Unzip(filepath.Join(s.dir, current), reader); err != nil {
		return err
	}
	db, err := createDb(s.dir)
	if err != nil {
		return err
	}
	s.db = db
	return nil

}

func (s *EventStateMachine) Close() error {
	s.db.Close()
	return nil
}

func createDb(dir string) (*gorocksdb.TransactionDB, error) {
	//启动时，如果快照存在，应该基于快照创建状态机
	snapshotDir := filepath.Join(dir, snapshot)
	currentDir := filepath.Join(dir, current)
	_, err := os.Stat(snapshotDir)
	if err == nil {
		if err := os.RemoveAll(currentDir); err != nil {
			return nil, err
		}
		if err := os.Rename(snapshotDir, currentDir); err != nil {
			return nil, err
		}
	}
	options := gorocksdb.NewDefaultOptions()
	options.SetCreateIfMissing(true)
	return gorocksdb.OpenTransactionDb(options, gorocksdb.NewDefaultTransactionDBOptions(), currentDir)
}
