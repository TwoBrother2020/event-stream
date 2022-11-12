package main

import (
	"github.com/hashicorp/raft"
	"github.com/tecbot/gorocksdb"
)

type partition struct {
	db           *gorocksdb.DB
	writeOptions *gorocksdb.WriteOptions
	readOptions  *gorocksdb.ReadOptions
	stableStore  *gorocksdb.ColumnFamilyHandle
}

func NewPartition(db *gorocksdb.DB, writeOptions *gorocksdb.WriteOptions, readOptions *gorocksdb.ReadOptions) *partition {
	columnFamily, err := db.CreateColumnFamily(gorocksdb.NewDefaultOptions(), "stable_store")
	if err != nil {
		return nil
	}
	return &partition{db: db, writeOptions: writeOptions, readOptions: readOptions, stableStore: columnFamily}
}

func (p *partition) Set(key []byte, val []byte) error {
	return p.db.PutCF(p.writeOptions, p.stableStore, key, val)
}

func (p *partition) Get(key []byte) ([]byte, error) {
	slice, err := p.db.GetCF(p.readOptions, p.stableStore, key)
	if err != nil {
		return nil, err
	}
	return slice.Data(), nil
}

func (p *partition) SetUint64(key []byte, val uint64) error {

	return p.Set(key, uint64ToBytes(val))
}

func (p *partition) GetUint64(key []byte) (uint64, error) {

	bytes, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	return bytesToUint64(bytes), nil

}

func (p *partition) FirstIndex() (uint64, error) {
	iterator := p.db.NewIterator(p.readOptions)
	iterator.SeekToFirst()
	key := iterator.Key()
	if len(key.Data()) == 0 {
		return uint64(0), nil
	}
	return bytesToUint64(key.Data()), nil
}

func (p *partition) LastIndex() (uint64, error) {

	iterator := p.db.NewIterator(p.readOptions)
	iterator.SeekToLast()
	key := iterator.Key()
	if len(key.Data()) == 0 {
		return uint64(0), nil
	}
	return bytesToUint64(key.Data()), nil
}

func (p *partition) GetLog(index uint64, log *raft.Log) error {
	val, err := p.db.Get(p.readOptions, uint64ToBytes(index))
	if err != nil {
		return err
	}
	if val == nil {
		return raft.ErrLogNotFound
	}
	return decodeMsgPack(val.Data(), log)
}

func (p *partition) StoreLog(log *raft.Log) error {
	return p.StoreLogs([]*raft.Log{log})
}

func (p *partition) StoreLogs(logs []*raft.Log) error {
	batch := gorocksdb.NewWriteBatch()
	for i := range logs {
		log := logs[i]
		buffer, err := encodeMsgPack(log)
		if err != nil {
			return err
		}
		batch.Put(uint64ToBytes(log.Index), buffer.Bytes())
	}
	err := p.db.Write(p.writeOptions, batch)
	return err
}

func (p *partition) DeleteRange(min, max uint64) error {
	err := p.db.DeleteFileInRange(gorocksdb.Range{Start: uint64ToBytes(min), Limit: uint64ToBytes(max)})
	return err
}
