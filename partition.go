package mm

import "github.com/coreos/etcd/raft"

type partition struct {
	nodeMap map[uint64]raft.Node
}
