package main

import (
	"github.com/hashicorp/raft"
	"github.com/stretchr/testify/assert"
	"github.com/tecbot/gorocksdb"
	"testing"
)

func TestPartition(t *testing.T) {
	dir := t.TempDir()
	db, err := createDb(dir)
	assert.Nil(t, err)

	partition := NewPartition(db, gorocksdb.NewDefaultWriteOptions(), gorocksdb.NewDefaultReadOptions())
	index, err := partition.FirstIndex()
	assert.Equal(t, uint64(0), index, "The two words should be the same.")
	latIndex, err := partition.LastIndex()
	assert.Equal(t, uint64(0), latIndex, "The two words should be the same.")
	log := &raft.Log{
		Index: 1,
		Data:  []byte("hello world"),
	}
	err = partition.StoreLog(log)
	assert.Nil(t, err)
	latIndex, err = partition.LastIndex()
	assert.Equal(t, uint64(1), latIndex, "The two words should be the same.")
	var data raft.Log
	err = partition.GetLog(1, &data)
	assert.Nil(t, err)
	assert.Equal(t, log.Index, data.Index)
	assert.Equal(t, log.Data, data.Data)

}

func createDb(dir string) (*gorocksdb.DB, error) {
	options := gorocksdb.NewDefaultOptions()
	options.SetCreateIfMissing(true)
	return gorocksdb.OpenDb(options, dir)
}
