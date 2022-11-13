package event_stream

import (
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/tecbot/gorocksdb"
	"io"
	"os"
)

type EventHandler struct {
	count     int
	dir       string
	db        *gorocksdb.TransactionDB
	options   *gorocksdb.Options
	txOptions *gorocksdb.TransactionDBOptions
}

func NewEventHandler(dir string, options *gorocksdb.Options, txOptions *gorocksdb.TransactionDBOptions) (*EventHandler, error) {
	db, err := gorocksdb.OpenTransactionDb(options, txOptions, dir)
	if err != nil {
		return nil, err
	}
	return &EventHandler{
		count:     0,
		dir:       dir,
		db:        db,
		options:   options,
		txOptions: txOptions,
	}, nil
}

func (p *EventHandler) Apply(log *raft.Log) interface{} {

	println("apply", log.Index)
	begin := p.db.TransactionBegin(gorocksdb.NewDefaultWriteOptions(), gorocksdb.NewDefaultTransactionOptions(), nil)

	begin.Put([]byte("consumer_index"), uint64ToBytes(log.Index))

	if err := begin.Commit(); err != nil {
		return err
	}

	return fmt.Sprintf("%s%d", "hello world", p.count)
}

func (p *EventHandler) Snapshot() (raft.FSMSnapshot, error) {
	println("create Snapshot")
	checkpoint, err := p.db.NewCheckpoint()
	if err != nil {
		return nil, err
	}
	return &CheckPointSnapshot{
		checkpoint: checkpoint,
	}, nil
}

func (p *EventHandler) Restore(snapshot io.ReadCloser) error {
	println("Restore Snapshot")

	p.db.Close()
	err := os.RemoveAll(p.dir)
	if err != nil {
		return err
	}
	err = unzip(p.dir, snapshot)
	if err != nil {
		return err
	}
	db, err := gorocksdb.OpenTransactionDb(gorocksdb.NewDefaultOptions(), gorocksdb.NewDefaultTransactionDBOptions(), p.dir)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

type CheckPointSnapshot struct {
	checkpoint *gorocksdb.Checkpoint
	dir        string
}

func (c *CheckPointSnapshot) Persist(sink raft.SnapshotSink) error {
	err := c.checkpoint.CreateCheckpoint(c.dir, 100)
	defer func() {
		if err != nil {
			sink.Cancel()
		}
	}()
	buffer, err := zipDir(c.dir)
	if err != nil {
		return err
	}
	sink.Write(buffer.Bytes())
	sink.Close()
	return nil

}

func (c *CheckPointSnapshot) Release() {
	c.checkpoint.Destroy()
}
