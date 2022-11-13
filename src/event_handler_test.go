package src

import (
	"github.com/hashicorp/raft"
	"github.com/stretchr/testify/assert"
	"github.com/tecbot/gorocksdb"
	"testing"
)

func TestEventHandler_Apply(t *testing.T) {
	options := gorocksdb.NewDefaultOptions()
	options.SetCreateIfMissing(true)
	dbOptions := gorocksdb.NewDefaultTransactionDBOptions()
	handler, err := NewEventHandler(t.TempDir(), options, dbOptions)
	assert.Nil(t, err)

	apply := handler.Apply(&raft.Log{
		Index: 0,
		Data:  []byte("hello wrold"),
	})

	//snapshot, err := handler.Snapshot()
	//assert.Nil(t, err)
	//sink := &raft.InmemSnapshotSink{}
	//follower, err := NewEventHandler(filepath.Join(t.TempDir(), "A"), options, dbOptions)
	println(apply)
}
