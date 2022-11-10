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
	//TODO implement me
	panic("implement me")
}

func (c cacheFSM) Snapshot() (raft.FSMSnapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (c cacheFSM) Restore(snapshot io.ReadCloser) error {
	//TODO implement me
	panic("implement me")
}

func Node(dir, name, host string, bootstrap bool) (*raft.Raft, error) {
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

	return newRaft, nil

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func TestAddNode(t *testing.T) {
	dir := "data"
	nodeA, err := Node(dir, "A", "localhost:5000", true)
	check(err)
	defer nodeA.Shutdown().Error()
	nodeB, err := Node(dir, "B", "localhost:50002", false)
	check(err)
	defer nodeB.Shutdown().Error()

	configurationFuture := nodeA.GetConfiguration()
	err = configurationFuture.Error()
	check(err)

	configuration := configurationFuture.Configuration()
	println(configuration.Servers)

	id, serverID := nodeA.LeaderWithID()
	println("leader", id, "address", serverID)

	time.Sleep(1 * time.Minute)
	//future := node.Apply([]byte("hello world"), 1*time.Minute)
	//err = future.Error()
	//check(err)

}
