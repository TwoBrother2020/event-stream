package key

import (
	"event-stream/protocol"
	"github.com/tecbot/gorocksdb"
)

type store struct {
	db *gorocksdb.DB
	// 数据存储目录
	dir        string
	jobHandle  *gorocksdb.ColumnFamilyHandle
	timeHandle *gorocksdb.ColumnFamilyHandle
}

func newStore(dir string) (*store, error) {
	options := gorocksdb.NewDefaultOptions()
	db, handles, err := gorocksdb.OpenDbColumnFamilies(options, dir, []string{"default", "job", "time"}, []*gorocksdb.Options{options, options, options})
	if err != nil {
		return nil, err
	}
	return &store{dir: dir, db: db, jobHandle: handles[1], timeHandle: handles[2]}, nil
}

func (s *store) Update(job protocol.JobCreate) {

	slice, err := s.db.GetCF(gorocksdb.NewDefaultReadOptions(), s.jobHandle, []byte(job.Name))
	if err != nil {
		return
	}

}
