package broker

import (
	"event-stream/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/lni/dragonboat/v4"
	"github.com/lni/dragonboat/v4/config"
	"github.com/lni/dragonboat/v4/raftio"
	sm "github.com/lni/dragonboat/v4/statemachine"
	"log"
	"path/filepath"
	"testing"
	"time"
)

type LeaderNotify struct {
	Notify  chan raftio.LeaderInfo
	Receive bool
}

func NewLeaderNotify() *LeaderNotify {
	return &LeaderNotify{Notify: make(chan raftio.LeaderInfo), Receive: true}
}

func (l *LeaderNotify) LeaderUpdated(info raftio.LeaderInfo) {
	if l.Receive {
		l.Notify <- info
	}

	println("不通知")
}

func TestDragonboat(t *testing.T) {

	dir := t.TempDir()
	leaderNotify := NewLeaderNotify()
	host1, err := createHost(filepath.Join(dir, "1"), "localhost:63001", leaderNotify)
	if err != nil {
		t.Error(err)
	}
	defer host1.Close()

	rc := config.Config{
		ReplicaID:          1,
		ShardID:            1,
		ElectionRTT:        10,
		HeartbeatRTT:       1,
		CheckQuorum:        true,
		SnapshotEntries:    10,
		CompactionOverhead: 5,
	}
	err = host1.StartOnDiskReplica(map[uint64]dragonboat.Target{1: "localhost:63001"}, false, func(shardID uint64, replicaID uint64) sm.IOnDiskStateMachine {

		return &EventStateMachine{
			shardID:   shardID,
			replicaID: replicaID,
			host:      host1,
		}
	}, rc)
	if err != nil {
		t.Error(err)
	}

	for {
		leaderInfo := <-leaderNotify.Notify
		if leaderInfo.LeaderID != 0 {
			leaderNotify.Receive = false
			println("receive", leaderInfo.ShardID, leaderInfo.ReplicaID, leaderInfo.LeaderID)
			break
		}
	}
	host2, err := createHost(filepath.Join(dir, "2"), "localhost:63002", nil)
	if err != nil {
		t.Error(err)
	}
	defer host2.Close()

	rc.ReplicaID = 2
	err = host2.StartOnDiskReplica(map[uint64]dragonboat.Target{},
		true,
		func(shardID uint64, replicaID uint64) sm.IOnDiskStateMachine {
			return &EventStateMachine{
				shardID:   shardID,
				replicaID: replicaID,
				host:      host1,
			}
		}, rc)
	if err != nil {
		t.Error(err)
	}

	id, _, b, err := host1.GetLeaderID(1)
	if b {
		println(id)
	}
	replica, err := host1.RequestAddReplica(1, 2, "localhost:63002", 0, 5*time.Second)
	if err != nil {
		return
	}
	result := <-replica.AppliedC()
	println(result.Completed())

	event := protocol.Event{
		Value: &protocol.Event_JobCreate{JobCreate: &protocol.JobCreate{Name: ""}},
	}
	marshal, err := proto.Marshal(&event)
	if err != nil {
		return
	}
	opSession := host1.GetNoOPSession(1)
	propose, err := host1.Propose(opSession, marshal, 5*time.Second)
	if err != nil {
		return
	}
	c := <-propose.AppliedC()
	if c.Committed() {
		println("")
	}

	time.Sleep(1 * time.Minute)
}

func createHost(dir, host string, notify *LeaderNotify) (*dragonboat.NodeHost, error) {
	nhc := config.NodeHostConfig{
		WALDir:         dir,
		NodeHostDir:    dir,
		RTTMillisecond: 200,
		RaftAddress:    host,
	}
	if notify != nil {
		nhc.RaftEventListener = notify
	}
	nh, err := dragonboat.NewNodeHost(nhc)
	if err != nil {
		return nil, err
	}
	return nh, nil
}

func TestEventStateMachine(t *testing.T) {
	event := &protocol.Event{
		Value: &protocol.Event_JobCompleted{},
	}
	event.GetValue()

	marshal, err := proto.Marshal(event)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	println(marshal)

}
