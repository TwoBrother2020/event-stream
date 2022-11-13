package src

import (
	"github.com/hashicorp/raft"
	"github.com/tecbot/gorocksdb"
)

type RocksdbStore struct {
	db           *gorocksdb.DB
	writeOptions *gorocksdb.WriteOptions
	readOptions  *gorocksdb.ReadOptions
	stableStore  *gorocksdb.ColumnFamilyHandle
}

func NewRocksdbStore(dir string) (*RocksdbStore, error) {

	options := gorocksdb.NewDefaultOptions()

	options.SetCreateIfMissing(true)
	options.SetCreateIfMissingColumnFamilies(true)

	db, familyHandles, err := gorocksdb.OpenDbColumnFamilies(options, dir, []string{"default", "stable_store"}, []*gorocksdb.Options{options, options})
	if err != nil {
		return nil, err
	}

	return &RocksdbStore{db: db, writeOptions: gorocksdb.NewDefaultWriteOptions(), readOptions: gorocksdb.NewDefaultReadOptions(), stableStore: familyHandles[1]}, nil
}

func (p *RocksdbStore) Set(key []byte, val []byte) error {
	return p.db.PutCF(p.writeOptions, p.stableStore, key, val)
}

func (p *RocksdbStore) Get(key []byte) ([]byte, error) {
	slice, err := p.db.GetCF(p.readOptions, p.stableStore, key)
	if err != nil {
		return nil, err
	}
	return slice.Data(), nil
}

func (p *RocksdbStore) SetUint64(key []byte, val uint64) error {

	return p.Set(key, uint64ToBytes(val))
}

func (p *RocksdbStore) GetUint64(key []byte) (uint64, error) {

	bytes, err := p.Get(key)
	if err != nil || len(bytes) == 0 {
		return 0, err
	}
	return bytesToUint64(bytes), nil

}

func (p *RocksdbStore) FirstIndex() (uint64, error) {
	iterator := p.db.NewIterator(p.readOptions)
	iterator.SeekToFirst()
	key := iterator.Key()
	if len(key.Data()) == 0 {
		return uint64(0), nil
	}
	return bytesToUint64(key.Data()), nil
}

func (p *RocksdbStore) LastIndex() (uint64, error) {

	iterator := p.db.NewIterator(p.readOptions)
	iterator.SeekToLast()
	key := iterator.Key()
	if len(key.Data()) == 0 {
		return uint64(0), nil
	}
	return bytesToUint64(key.Data()), nil
}

func (p *RocksdbStore) GetLog(index uint64, log *raft.Log) error {
	val, err := p.db.Get(p.readOptions, uint64ToBytes(index))
	if err != nil {
		return err
	}
	if val == nil {
		return raft.ErrLogNotFound
	}
	return decodeMsgPack(val.Data(), log)
}

func (p *RocksdbStore) StoreLog(log *raft.Log) error {
	return p.StoreLogs([]*raft.Log{log})
}

func (p *RocksdbStore) StoreLogs(logs []*raft.Log) error {
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

func (p *RocksdbStore) DeleteRange(min, max uint64) error {
	err := p.db.DeleteFileInRange(gorocksdb.Range{Start: uint64ToBytes(min), Limit: uint64ToBytes(max)})
	return err
}
