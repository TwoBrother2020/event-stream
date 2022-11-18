package broker

import (
	sm "github.com/lni/dragonboat/v4/statemachine"
	"io"
)

type EventStateMachine struct {
	clusterID uint64
	nodeID    uint64
}

func (s *EventStateMachine) Open(stopc <-chan struct{}) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *EventStateMachine) Update(entries []sm.Entry) ([]sm.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (s *EventStateMachine) Lookup(i interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s *EventStateMachine) Sync() error {
	//TODO implement me
	panic("implement me")
}

func (s *EventStateMachine) PrepareSnapshot() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s *EventStateMachine) SaveSnapshot(i interface{}, writer io.Writer, i2 <-chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventStateMachine) RecoverFromSnapshot(reader io.Reader, i <-chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventStateMachine) Close() error {
	//TODO implement me
	panic("implement me")
}
