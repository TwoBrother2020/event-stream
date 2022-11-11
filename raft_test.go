package main

import (
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb/v2"
	"io"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type cacheFSM struct {
}

func newCacheFSM() *cacheFSM {
	return &cacheFSM{}
}

func (c cacheFSM) Apply(log *raft.Log) interface{} {
	s := ""
	return s
}

func (c cacheFSM) Snapshot() (raft.FSMSnapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (c cacheFSM) Restore(snapshot io.ReadCloser) error {
	//TODO implement me
	panic("implement me")
}

func Node(dir, name, host string, bootstrap bool) (*raft.Raft, *raftboltdb.BoltStore, error) {
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(name)
	nodeDir := filepath.Join(dir, name)
	_, err := os.Stat(nodeDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(nodeDir, 0755)
	}
	check(err)
	logStore, err := raftboltdb.NewBoltStore(filepath.Join(nodeDir, "raft-log.bolt"))
	check(err)

	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(nodeDir, "raft-stable.bolt"))
	check(err)

	snapshotStore, err := raft.NewFileSnapshotStore(nodeDir, 1, os.Stderr)
	check(err)

	address, err := net.ResolveTCPAddr("tcp", host)
	check(err)

	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	check(err)

	newRaft, err := raft.NewRaft(raftConfig, newCacheFSM(), logStore, stableStore, snapshotStore, transport)
	check(err)

	if bootstrap {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      raft.ServerID(name),
					Address: raft.ServerAddress(host),
				},
			},
		}
		newRaft.BootstrapCluster(configuration)
	}

	return newRaft, logStore, nil

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func BenchmarkNodeWrite(b *testing.B) {

	dir := "data"
	nodeA, logA, err := Node(dir, "A", "localhost:5000", false)
	check(err)
	defer nodeA.Shutdown()
	nodeB, _, err := Node(dir, "B", "localhost:50002", false)
	check(err)
	defer nodeB.Shutdown()

	configurationFuture := nodeA.GetConfiguration()
	err = configurationFuture.Error()
	check(err)

	configuration := configurationFuture.Configuration()
	println(configuration.Servers)
	leader := <-nodeA.LeaderCh()
	if leader {
		index, err := logA.LastIndex()
		check(err)
		println("last", index)
		for i := 0; i < b.N; i++ {
			future := nodeA.Apply([]byte("hello world2"), 1*time.Minute)

			if err := future.Error(); err != nil {
				println("error")
				return
			}
		}
	}

}
