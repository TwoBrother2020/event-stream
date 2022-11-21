package broker

import (
	"github.com/stretchr/testify/assert"
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
