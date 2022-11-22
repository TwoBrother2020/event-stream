package key

import (
	"errors"
	"event-stream/protocol"
	"github.com/tecbot/gorocksdb"
	"google.golang.org/protobuf/proto"
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

func (s *store) Create(job protocol.JobCreate) error {

	slice, err := s.db.GetCF(gorocksdb.NewDefaultReadOptions(), s.jobHandle, []byte(job.Name))
	if err != nil {
		return err
	}
	if len(slice.Data()) > 0 {
		return errors.New("job name exist")
	}

	value, err := proto.Marshal(&job)
	if err != nil {
		return err
	}
	err = s.db.PutCF(gorocksdb.NewDefaultWriteOptions(), s.jobHandle, []byte(job.Name), value)
	if err != nil {
		return err
	}

	return nil
}

//func (s *store) Update(job protocol.JobCreate) error {
//
//	slice, err := s.db.GetCF(gorocksdb.NewDefaultReadOptions(), s.jobHandle, []byte(job.Name))
//	if err != nil {
//		return err
//	}
//	if len(slice.Data()) == 0 {
//		s.db.PutCF(gorocksdb.NewDefaultWriteOptions(), s.jobHandle, []byte(job.Name))
//	}
//
//	return nil
//}
//
//func (s *store) Cancel(job protocol.JobCreate) error {
//
//	slice, err := s.db.GetCF(gorocksdb.NewDefaultReadOptions(), s.jobHandle, []byte(job.Name))
//	if err != nil {
//		return err
//	}
//	if len(slice.Data()) == 0 {
//		s.db.PutCF(gorocksdb.NewDefaultWriteOptions(), s.jobHandle, []byte(job.Name))
//	}
//
//	return nil
//}
//
//func (s *store) Delete(job protocol.JobCreate) error {
//
//	slice, err := s.db.GetCF(gorocksdb.NewDefaultReadOptions(), s.jobHandle, []byte(job.Name))
//	if err != nil {
//		return err
//	}
//	if len(slice.Data()) == 0 {
//		s.db.PutCF(gorocksdb.NewDefaultWriteOptions(), s.jobHandle, []byte(job.Name))
//	}
//	return nil
//}
