package src

import (
	"github.com/cockroachdb/pebble"
	"github.com/hashicorp/raft"
	"github.com/stretchr/testify/assert"
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
	handler, err := NewEventHandler(dir, gorocksdb.NewDefaultOptions(), gorocksdb.NewDefaultTransactionDBOptions())
	check(err)
	newRaft, err := raft.NewRaft(raftConfig, handler, store, store, snapshotStore, transport)
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

func TestRocks(t *testing.T) {
	key := []byte("1")
	value := []byte("hello world")
	dir := t.TempDir()
	options := gorocksdb.NewDefaultOptions()
	options.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(options, filepath.Join(dir, "A"))
	assert.Nil(t, err)
	snapshot := db.NewSnapshot()

	readOptions := gorocksdb.NewDefaultReadOptions()
	readOptions.SetSnapshot(snapshot)

	err = db.Put(gorocksdb.NewDefaultWriteOptions(), key, value)
	assert.Nil(t, err)
	slice, err := db.Get(gorocksdb.NewDefaultReadOptions(), key)
	assert.Nil(t, err)
	println("data", string(slice.Data()))
	slice, err = db.Get(readOptions, key)
	assert.Nil(t, err)
	println("data", string(slice.Data()))
	db.ReleaseSnapshot(snapshot)
	checkpoint, err := db.NewCheckpoint()
	assert.Nil(t, err)
	err = checkpoint.CreateCheckpoint(filepath.Join(dir, "B"), 100)
	assert.Nil(t, err)
	b, err := gorocksdb.OpenDb(gorocksdb.NewDefaultOptions(), filepath.Join(dir, "B"))
	assert.Nil(t, err)
	slice, err = b.Get(gorocksdb.NewDefaultReadOptions(), key)

	assert.Nil(t, err)
	println("data-b", string(slice.Data()))

}
