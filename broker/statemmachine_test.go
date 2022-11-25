package key

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/dgraph-io/badger/v3"
	"github.com/stretchr/testify/assert"
	"log"
	"path/filepath"
	"testing"
)

// BenchmarkBadger-8         105655             10552 ns/op
func BenchmarkBadger(b *testing.B) {

	db, err := badger.Open(badger.DefaultOptions(b.TempDir()))
	db.MaxVersion()
	assert.Nil(b, err)
	defer db.Close()
	for i := 0; i < b.N; i++ {
		txn := db.NewTransaction(true)
		bytes := []byte(fmt.Sprintf("hello %d", i))
		err := txn.Set(bytes, bytes)
		assert.Nil(b, err)
		err = txn.Commit()
	}
}

//BenchmarkRocksDB-8        349010              3369 ns/op

// BenchmarkPebble-8         870882              1202 ns/op
func BenchmarkPebble(b *testing.B) {
	dir := b.TempDir()
	db, err := pebble.Open(dir, &pebble.Options{})
	assert.Nil(b, err)
	defer db.Close()
	err = db.Checkpoint("", nil)
	for i := 0; i < b.N; i++ {
		bytes := []byte(fmt.Sprintf("hello %d", i))
		err := db.Set(bytes, bytes, pebble.NoSync)
		assert.Nil(b, err)
	}
}

func TestPebble(t *testing.T) {
	tempDir := t.TempDir()
	db, err := pebble.Open(tempDir, &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}
	key := []byte("hello")
	if err := db.Set(key, []byte("world"), pebble.Sync); err != nil {
		log.Fatal(err)
	}

	dir := filepath.Join(tempDir, "checkpoint")
	err = db.Checkpoint(dir)
	assert.Nil(t, err)
	value, closer, err := db.Get(key)
	assert.Nil(t, err)

	fmt.Printf("%s %s\n", key, value)
	closer.Close()
	err = db.Delete(key, pebble.Sync)
	assert.Nil(t, err)
	db2, err := pebble.Open(dir, &pebble.Options{})
	assert.Nil(t, err)
	value2, closer2, err2 := db2.Get(key)
	assert.Nil(t, err2)
	closer2.Close()
	fmt.Printf("%s %s\n", key, value2)

}

func TestEventStateMachine_Open(t *testing.T) {
	dir := t.TempDir()

	stateMachine := NewEventStateMachine(dir, 1, 1)
	c := make(chan struct{})
	open, err := stateMachine.Open(c)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), open)

}
