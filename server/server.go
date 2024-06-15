package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/logger"
	"github.com/lni/goutils/syncutil"
)

func main() {
	nodeID := flag.Int("nodeid", 1, "Server's NodeID")
	addr := flag.String("address", "", "Address where the server will run. Should be in the format: address:port")
	join := flag.Bool("join", false, "If the server is joining an existing cluster")
	clusterID := flag.Int(
		"clusterid",
		128,
		"ID of the cluster which the server wants to join/create. If there's already a cluster running with the same ID, you must use the -join flag (default 128)",
	)
	flag.Parse()

	if len(*addr) == 0 {
		fmt.Fprintf(os.Stderr, "address not specified")
		os.Exit(1)
	}

	initialMembers := make(map[uint64]string)

	if !*join {
		initialMembers[uint64(*nodeID)] = *addr
	}

	logger.GetLogger("raft").SetLevel(logger.ERROR)
	logger.GetLogger("rsm").SetLevel(logger.WARNING)
	logger.GetLogger("transport").SetLevel(logger.ERROR)
	logger.GetLogger("grpc").SetLevel(logger.WARNING)

	cfg := config.Config{
		NodeID:             uint64(*nodeID),
		ClusterID:          uint64(*clusterID),
		ElectionRTT:        10,
		HeartbeatRTT:       1,
		CheckQuorum:        true,
		SnapshotEntries:    100,
		CompactionOverhead: 5,
	}

	dataDir := filepath.Join(
		"tuple-space-storage",
		fmt.Sprintf("node-%d", *nodeID),
	)

	nodeHostCfg := config.NodeHostConfig{
		WALDir:         dataDir,
		NodeHostDir:    dataDir,
		RTTMillisecond: 200,
		RaftAddress:    *addr,
	}

	nodeHost, err := dragonboat.NewNodeHost(nodeHostCfg)
	if err != nil {
		panic(err)
	}

	if err := nodeHost.StartCluster(initialMembers, *join, nil, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start cluster %d: %v", *clusterID, err)
		os.Exit(1)
	}
}
