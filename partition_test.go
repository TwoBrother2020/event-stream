package mm

import (
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"testing"
	"time"
)

func TestStartNode(t *testing.T) {
	storage := raft.NewMemoryStorage()
	config := &raft.Config{
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256}
	node := raft.StartNode(config, nil)

	ticker := time.NewTicker(3 * time.Second).C
	for {
		select {
		case <-ticker:
			node.Tick()
		case rd := <-node.Ready():
			storage.SetHardState(rd.HardState)
			storage.Append(rd.Entries)
			//send(rd.Messages)
			if !raft.IsEmptySnap(rd.Snapshot) {
				storage.ApplySnapshot(rd.Snapshot)
			}
			for _, entry := range rd.CommittedEntries {
				//process(entry)
				if entry.Type == raftpb.EntryConfChange {
					var cc raftpb.ConfChange
					cc.Unmarshal(entry.Data)
					node.ApplyConfChange(cc)
				}
			}
			node.Advance()
		}
	}
}
