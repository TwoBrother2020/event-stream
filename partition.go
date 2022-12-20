package mm

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/raft/v3"
	"go.etcd.io/etcd/raft/v3/raftpb"
	"go.etcd.io/etcd/server/v3/etcdserver/api/snap"
	"go.etcd.io/etcd/server/v3/wal"
	"time"
)

type commit struct {
	data       []string
	applyDoneC chan<- struct{}
}

// partition
type partition struct {
	id      int      // client ID for raft session
	peers   []string // raft peer URLs
	join    bool     // node is joining an existing cluster
	walDir  string   // path to WAL directory
	snapDir string   // path to snapshot directory

	node          raft.Node
	raftStorage   *raft.MemoryStorage
	wal           *wal.WAL
	confState     raftpb.ConfState
	snapshotIndex uint64
	appliedIndex  uint64

	proposeC    <-chan string            // proposed messages (k,v)
	confChangeC <-chan raftpb.ConfChange // proposed cluster config changes
	commitC     chan<- *commit           // entries committed to log (k,v)
	errorC      chan<- error             // errors from raft session

	stopc            chan struct{} // signals proposal channel closed
	snapshotterReady chan *snap.Snapshotter
}

func newPartition(id int, peers []string, join bool, getSnapshot func() ([]byte, error), proposeC <-chan string,
	confChangeC <-chan raftpb.ConfChange) (<-chan *commit, <-chan error, <-chan *snap.Snapshotter) {
	commitC := make(chan *commit)
	errorC := make(chan error)

	p := &partition{
		proposeC:    proposeC,
		confChangeC: confChangeC,
		commitC:     commitC,
		errorC:      errorC,
		id:          id,
		peers:       peers,
		join:        join,
		walDir:      fmt.Sprintf("partition-%d", id),
		snapDir:     fmt.Sprintf("partition-%d-snap", id),
		stopc:       make(chan struct{}),

		snapshotterReady: make(chan *snap.Snapshotter, 1),
		// rest of structure populated after WAL replay
	}
	return commitC, errorC, p.snapshotterReady

}

func (p *partition) start() {

	snap, err := p.raftStorage.Snapshot()
	if err != nil {
		panic(err)
	}
	p.confState = snap.Metadata.ConfState
	p.snapshotIndex = snap.Metadata.Index
	p.appliedIndex = snap.Metadata.Index

	defer p.wal.Close()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// send proposals over raft
	go func() {
		confChangeCount := uint64(0)

		for p.proposeC != nil && p.confChangeC != nil {
			select {
			case prop, ok := <-p.proposeC:
				if !ok {
					p.proposeC = nil
				} else {
					// blocks until accepted by raft state machine
					p.node.Propose(context.TODO(), []byte(prop))
				}

			case cc, ok := <-p.confChangeC:
				if !ok {
					p.confChangeC = nil
				} else {
					confChangeCount++
					cc.ID = confChangeCount
					p.node.ProposeConfChange(context.TODO(), cc)
				}
			}
		}
		// client closed channel; shutdown raft if not already
		close(p.stopc)
	}()

	// event loop on raft state machine updates

	for {
		select {
		case <-ticker.C:
			p.node.Tick()
			// store raft entries to wal, then publish over commit channel
		case rd := <-p.node.Ready():
			p.saveToStorage(rd.HardState, rd.Entries)
			p.send(rd.Messages)
			if !raft.IsEmptySnap(rd.Snapshot) {
				processSnapshot(rd.Snapshot)
			}
			for _, entry := range rd.CommittedEntries {
				process(entry)
				if entry.Type == raftpb.EntryConfChange {
					var cc raftpb.ConfChange
					cc.Unmarshal(entry.Data)
					p.node.ApplyConfChange(cc)
				}
			}
			p.node.Advance()
		case <-p.stopc:
			return
		}
	}
}

func process(entry raftpb.Entry) {

}

func processSnapshot(snapshot raftpb.Snapshot) {

}

func (p *partition) send(messages []raftpb.Message) {

}

func (p *partition) saveToStorage(state raftpb.HardState, entries []raftpb.Entry) error {
	return p.wal.Save(state, entries)
}
