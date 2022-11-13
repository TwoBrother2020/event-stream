package src

import (
	"github.com/cockroachdb/pebble"
	"github.com/hashicorp/raft"
	"github.com/tecbot/gorocksdb"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func Node(dir, name, host string, bootstrap bool) (*raft.Raft, *RocksdbStore, error) {
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(name)
	nodeDir := filepath.Join(dir, name)
	_, err := os.Stat(nodeDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(nodeDir, 0755)
	}
	check(err)

	store, _ := NewRocksdbStore(nodeDir)

	snapshotStore, err := raft.NewFileSnapshotStore(nodeDir, 1, os.Stderr)
	check(err)

	address, err := net.ResolveTCPAddr("tcp", host)
	check(err)

	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	check(err)

	newRaft, err := raft.NewRaft(raftConfig, NewEventHandler(), store, store, snapshotStore, transport)
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

	return newRaft, store, nil

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestNodeWrite(t *testing.T) {

	dir := "data"
	nodeA, logA, err := Node(dir, "A", "localhost:5000", true)
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
		var data string
		for i := 0; i < 100; i++ {
			index, err := logA.LastIndex()
			check(err)
			println("last index", index)
			if len(data) == 0 {
				data = "hello world"
			}
			future := nodeA.Apply([]byte(data), 1*time.Minute)
			if err := future.Error(); err != nil {
				return
			}
			data = future.Response().(string)

			println("return ", data)
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
