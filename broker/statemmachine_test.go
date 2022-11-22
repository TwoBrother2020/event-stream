package key

import (
	"github.com/stretchr/testify/assert"
	"github.com/tecbot/gorocksdb"
	"testing"
)

func TestEventStateMachine_Open(t *testing.T) {
	dir := t.TempDir()

	stateMachine := NewEventStateMachine(dir, 1, 1)
	c := make(chan struct{})
	open, err := stateMachine.Open(c)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), open)

}

func TestRocksDBTransaction(t *testing.T) {
	dir := t.TempDir()
	options := gorocksdb.NewDefaultOptions()
	options.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(options, dir)
	readOptions := gorocksdb.NewDefaultReadOptions()
	readOptions.SetPrefixSameAsStart(true)
	assert.Nil(t, err)
	transactionDb, err := gorocksdb.OpenTransactionDb(gorocksdb.NewDefaultOptions(), gorocksdb.NewDefaultTransactionDBOptions(), dir)
	assert.Nil(t, err)

	defer db.Close()
	defer transactionDb.Close()
}
