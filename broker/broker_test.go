package key

import (
	"github.com/lni/dragonboat/v4"
	"testing"
	"time"
)

func TestBroker_LeaderUpdated(t *testing.T) {
	broker1 := NewBroker(t.TempDir(), "localhost:63001", map[uint64]dragonboat.Target{1: "localhost:63001", 2: "localhost:63002"}, 1)

	err := broker1.run()
	if err != nil {
		t.Fatal(err)
	}
	broker2 := NewBroker(t.TempDir(), "localhost:63002", map[uint64]dragonboat.Target{1: "localhost:63001", 2: "localhost:63002"}, 2)

	err = broker2.run()
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(15 * time.Second)

	_, err = broker2.Propose(1, []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(15 * time.Minute)

}
