package key

import (
	"event-stream/util"
	sm "github.com/lni/dragonboat/v4/statemachine"
	"io"
	"path/filepath"
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
	responses     chan *Response
	// 数据存储目录
	dir    string
	leader bool
}

func (s *EventStateMachine) Lookup(i interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func NewEventStateMachine(dir string, shardID uint64, replicaID uint64) *EventStateMachine {
	return &EventStateMachine{dir: dir, shardID: shardID, replicaID: replicaID}
}

func (s *EventStateMachine) Open(stopc <-chan struct{}) (uint64, error) {
	return 0, nil

}

func (s *EventStateMachine) Update(entries []sm.Entry) (res []sm.Entry, err error) {

	for i := range entries {
		entry := entries[i]
		println(string(entry.Cmd))
		s.responses <- &Response{ShardID: s.shardID, data: &sm.Result{Value: entry.Index, Data: entry.Cmd}}
	}

	return entries, nil
}

func (s *EventStateMachine) Sync() error {
	return nil
}

func (s *EventStateMachine) PrepareSnapshot() (interface{}, error) {

	return nil, nil

}

func (s *EventStateMachine) SaveSnapshot(point interface{}, writer io.Writer, notify <-chan struct{}) error {
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

	if err := util.Unzip(filepath.Join(s.dir, current), reader); err != nil {
		return err
	}
	return nil

}

func (s *EventStateMachine) Close() error {
	return nil
}
