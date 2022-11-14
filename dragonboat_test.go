package main

import (
	"fmt"
	"github.com/lni/dragonboat/v4"
	"github.com/lni/dragonboat/v4/config"
	"os"
	"path/filepath"
	"testing"
)

func TestDragonboat(t *testing.T) {

	datadir := filepath.Join(t.TempDir(), fmt.Sprintf("node%d", 1))
	nhc := config.NodeHostConfig{
		WALDir:         datadir,
		NodeHostDir:    datadir,
		RTTMillisecond: 200,
		RaftAddress:    "localhost:63001",
	}
	nh, err := dragonboat.NewNodeHost(nhc)
	if err != nil {
		t.Error(err)
	}
	defer nh.Close()

	rc := config.Config{
		ReplicaID:          uint64(1),
		ShardID:            exampleShardID,
		ElectionRTT:        10,
		HeartbeatRTT:       1,
		CheckQuorum:        true,
		SnapshotEntries:    10,
		CompactionOverhead: 5,
	}
	if err := nh.StartOnDiskReplica(map[uint64]dragonboat.Target{1: "localhost:63001"}, false, NewDiskKV, rc); err != nil {
		fmt.Fprintf(os.Stderr, "failed to add cluster, %v\n", err)
		t.Error(err)
	}

}
