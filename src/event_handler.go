package src

import (
	"fmt"
	"github.com/hashicorp/raft"
	"io"
)

type EventHandler struct {
	count int
}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (p *EventHandler) Apply(log *raft.Log) interface{} {
	data := log.Data
	p.count++
	return fmt.Sprintf("%s%d", string(data), p.count)
}

func (p *EventHandler) Snapshot() (raft.FSMSnapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (p *EventHandler) Restore(snapshot io.ReadCloser) error {
	//TODO implement me
	panic("implement me")
}
