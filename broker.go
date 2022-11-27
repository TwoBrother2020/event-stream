package mm

import (
	"github.com/lni/dragonboat/v4"
	"github.com/lni/dragonboat/v4/config"
	"github.com/lni/dragonboat/v4/raftio"
	sm "github.com/lni/dragonboat/v4/statemachine"
	"log"
	"os"
	"time"
)

type Broker struct {
	dir          string
	host         string
	node         *dragonboat.NodeHost
	targets      map[uint64]dragonboat.Target
	replicaID    uint64
	stateMachine map[uint64]map[uint64]*DiskKV
	signal       chan os.Signal
}

func NewBroker(dir string, host string, targets map[uint64]dragonboat.Target, replicaID uint64) *Broker {
	return &Broker{dir: dir, host: host, targets: targets, replicaID: replicaID, stateMachine: map[uint64]map[uint64]*DiskKV{}}
}

func (b *Broker) LeaderUpdated(info raftio.LeaderInfo) {
	if replicaMap, ok := b.stateMachine[info.ShardID]; ok {
		stateMachine := replicaMap[info.ReplicaID]
		if stateMachine.replicaID == info.LeaderID {
			log.Printf("%s ShardID %d ReplicaID %d  became leader", b.host, info.ShardID, info.ReplicaID)
		} else {
		}
	}
}

func (b *Broker) createStateMachine(shardID uint64, replicaID uint64) sm.IOnDiskStateMachine {
	state := &DiskKV{
		shardID:   shardID,
		replicaID: replicaID,
		nh:        b.node,
	}
	if replicaMap, ok := b.stateMachine[shardID]; ok {
		replicaMap[replicaID] = state
	} else {
		b.stateMachine[shardID] = map[uint64]*DiskKV{replicaID: state}
	}
	return state
}

func (b *Broker) run() error {

	nhc := config.NodeHostConfig{
		WALDir:         b.dir,
		NodeHostDir:    b.dir,
		RTTMillisecond: 200,
		RaftAddress:    b.host,
	}
	nhc.RaftEventListener = b
	nh, err := dragonboat.NewNodeHost(nhc)
	if err != nil {
		return err
	}
	b.node = nh
	rc := config.Config{
		ReplicaID:          0,
		ShardID:            1,
		ElectionRTT:        10,
		HeartbeatRTT:       1,
		CheckQuorum:        true,
		SnapshotEntries:    10,
		CompactionOverhead: 5,
	}
	rc.ReplicaID = b.replicaID
	if nh.StartOnDiskReplica(b.targets, false, b.createStateMachine, rc); err != nil {
		return err
	}
	return nil
}

func (b *Broker) Propose(shardId uint64, cmd []byte) (*dragonboat.RequestState, error) {

	noOPSession := b.node.GetNoOPSession(shardId)
	b.node.GetLeaderID(shardId)
	return b.node.Propose(noOPSession, cmd, 30*time.Second)

}

func (b *Broker) close() {
	if b.node != nil {
		b.node.Close()
	}
}
