package main

import (
	"github.com/cockroachdb/pebble"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb/v2"
	"github.com/tecbot/gorocksdb"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
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

func TestNodeWrite(t *testing.T) {

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
		future := nodeA.Apply([]byte("hello world2"), 1*time.Minute)

		if err := future.Error(); err != nil {
			println("error")
			return
		}
	}

}

func BenchmarkPebbleWrite(b *testing.B) {
	dir := b.TempDir()
	db, err := pebble.Open(dir, &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		keyStr := "test" + strconv.Itoa(i)
		var key = []byte(keyStr)
		if err := db.Set(key, key, pebble.NoSync); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkRocksDbWrite(b *testing.B) {

	dir := b.TempDir()
	options := gorocksdb.NewDefaultOptions()
	db, err := gorocksdb.OpenDb(options, dir)
	if err != nil {
		return
	}
	defer db.Close()
	writeOptions := gorocksdb.NewDefaultWriteOptions()
	for i := 0; i < b.N; i++ {
		keyStr := "test" + strconv.Itoa(i)
		var key = []byte(keyStr)
		err := db.Put(writeOptions, key, key)
		if err != nil {
			b.Fatal(err)
		}
	}
}
