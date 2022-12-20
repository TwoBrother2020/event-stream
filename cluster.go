package mm

type ClusterConfig struct {
	partitions []PartitionConfig
}
type PartitionConfig struct {
	id      int
	host    string
	remotes []string
}

type Cluster struct {
}

func NewCluster(config *ClusterConfig) *Cluster {
	return &Cluster{}
}
