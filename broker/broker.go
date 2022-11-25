package key

import (
	"github.com/lni/dragonboat/v4"
	"github.com/lni/dragonboat/v4/config"
	"github.com/lni/dragonboat/v4/raftio"
	sm "github.com/lni/dragonboat/v4/statemachine"
	"log"
	"os"
	"time"
)

type Response struct {
	ShardID uint64
	data    *sm.Result
}

type Broker struct {
	dir          string
	host         string
	node         *dragonboat.NodeHost
	targets      map[uint64]dragonboat.Target
	replicaID    uint64
	stateMachine map[uint64]map[uint64]*EventStateMachine
	response     chan *Response
	signal       chan os.Signal
}

func NewBroker(dir string, host string, targets map[uint64]dragonboat.Target, replicaID uint64) *Broker {
	return &Broker{dir: dir, host: host, targets: targets, replicaID: replicaID, response: make(chan *Response, 100), stateMachine: map[uint64]map[uint64]*EventStateMachine{}}
}

func (b *Broker) LeaderUpdated(info raftio.LeaderInfo) {
	if replicaMap, ok := b.stateMachine[info.ShardID]; ok {
		stateMachine := replicaMap[info.ReplicaID]
		if stateMachine.replicaID == info.LeaderID {
			log.Printf("%s ShardID %d ReplicaID %d  became leader", b.host, info.ShardID, info.ReplicaID)
			stateMachine.leader = true

		} else {
			stateMachine.leader = false
		}
	}
}

func (b *Broker) createStateMachine(shardID uint64, replicaID uint64) sm.IOnDiskStateMachine {
	state := &EventStateMachine{
		shardID:   shardID,
		replicaID: replicaID,
		responses: b.response,
	}
	if replicaMap, ok := b.stateMachine[shardID]; ok {
		replicaMap[replicaID] = state
	} else {
		b.stateMachine[shardID] = map[uint64]*EventStateMachine{replicaID: state}
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
	go b.writeResponse()
	return nil
}

func (b *Broker) writeResponse() {
	for {
		select {
		case res := <-b.response:
			log.Printf("%s write shardId %d index %d %s \n", b.host, res.ShardID, res.data.Value, string(res.data.Data))
			noOPSession := b.node.GetNoOPSession(res.ShardID)
			_, err := b.node.Propose(noOPSession, res.data.Data, 30*time.Second)
			if err != nil {
				log.Printf("writeResponse error %s", err.Error())
				continue
			}
		case <-b.signal:
			return
		}
	}
}

func (b *Broker) Propose(shardId uint64, cmd []byte) (*dragonboat.RequestState, error) {

	noOPSession := b.node.GetNoOPSession(shardId)
	return b.node.Propose(noOPSession, cmd, 30*time.Second)

}

func (b *Broker) close() {
	if b.node != nil {
		b.node.Close()
	}
}
