package broker

import (
	"fmt"
	"github.com/lni/dragonboat/v4"
	"github.com/lni/dragonboat/v4/config"
	"path/filepath"
	"testing"
	"time"
)

func TestDragonboat(t *testing.T) {

	dir := t.TempDir()
	node1, err := creatNode(dir, "localhost:63001", 1)
	if err != nil {
		t.Error(err)
	}
	defer node1.Close()

	rc := config.Config{
		ReplicaID:          1,
		ShardID:            1,
		ElectionRTT:        10,
		HeartbeatRTT:       1,
		CheckQuorum:        true,
		SnapshotEntries:    10,
		CompactionOverhead: 5,
	}
	err = node1.StartOnDiskReplica(map[uint64]dragonboat.Target{1: "localhost:63001"}, false, NewDiskKV, rc)
	if err != nil {
		t.Error(err)
	}

	node2, err := creatNode(dir, "localhost:63002", 2)
	if err != nil {
		t.Error(err)
	}
	defer node2.Close()
	rc.ReplicaID = 2

	time.Sleep(1 * time.Minute)
}

func creatNode(dir, host string, node uint64) (*dragonboat.NodeHost, error) {
	datadir := filepath.Join(dir, fmt.Sprintf("node%d", node))
	nhc := config.NodeHostConfig{
		WALDir:         datadir,
		NodeHostDir:    datadir,
		RTTMillisecond: 200,
		RaftAddress:    host,
	}
	nh, err := dragonboat.NewNodeHost(nhc)
	if err != nil {
		return nil, err
	}

	return nh, nil
}

func TestCar(t *testing.T) {

}
