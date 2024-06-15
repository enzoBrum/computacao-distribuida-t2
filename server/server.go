package main

import (
	tuple_spaces "computacao-distribuida/tuple-spaces"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/logger"
	"google.golang.org/protobuf/proto"
)

// global vars, global vars everywhere
var nodeHost *dragonboat.NodeHost

const clusterID uint64 = 128

func handleRead(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to read tuple\n")
	var str_tuple []string
	if err := json.NewDecoder(r.Body).Decode(&str_tuple); err != nil {
		fmt.Printf("Error while deserializing tuple %v\n", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	tuple := tuple_spaces.Tuple{Values: str_tuple}
	fmt.Printf("Reading tuple: %v\n", &tuple)

	for i := 0; i < 10; i += 1 {
		fmt.Printf("Trying to make request: %d/10\n", i+1)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		result, err := nodeHost.SyncRead(ctx, clusterID, &tuple)
		foundTuple, ok := result.(*tuple_spaces.Tuple)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Errror: %v\b", err)
			return
		}
		if !ok {
			fmt.Println("Lookup did not return *Tuple")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err == nil {
			jsonTuple, _ := json.Marshal(foundTuple.Values)
			w.Write(jsonTuple)
			return
		}
		if err != dragonboat.ErrTimeout {
			fmt.Printf("Error while making request: %v", err)
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		}
	}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to add tuple\n")
	var str_tuple []string
	var buf []byte
	r.Body.Read(buf)
	fmt.Printf("ABCD: %d\n", len(buf))
	fmt.Printf("ABCD: %s\n", string(buf))
	if err := json.NewDecoder(r.Body).Decode(&str_tuple); err != nil {
		fmt.Printf("Error while deserializing tuple %v\n", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tuple := tuple_spaces.Tuple{Values: str_tuple}
	fmt.Printf("Adding tuple: %v\n", &tuple)
	req := tuple_spaces.RequestData{Cmd: tuple_spaces.Command_WRITE, Tuple: &tuple}
	serialized, err := proto.Marshal(&req)

	if err != nil {
		fmt.Printf("Error while serializing request: %v\n", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	session, err := nodeHost.SyncGetSession(ctx, clusterID)
	if err != nil {
		fmt.Printf("Error while creating raft client session: %v\n", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for i := 0; i < 10; i += 1 {
		fmt.Printf("Trying to make request: %d/10\n", i+1)
		_, err := nodeHost.SyncPropose(ctx, session, serialized)
		if err == nil {
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Errror: %v\b", err)
			return
		}
		if err != dragonboat.ErrTimeout {
			fmt.Printf("Error while making request: %v\n", err)
			http.Error(w, fmt.Sprintf("Error: %v\n", err), http.StatusInternalServerError)
		}
		ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
	}
}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to get tuple\n")
	var str_tuple []string
	if err := json.NewDecoder(r.Body).Decode(&str_tuple); err != nil {
		fmt.Printf("Error while deserializing tuple %v\n", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tuple := tuple_spaces.Tuple{Values: str_tuple}
	fmt.Printf("Reading tuple: %v\n", &tuple)

	session, err := nodeHost.SyncGetSession(ctx, clusterID)
	if err != nil {
		fmt.Printf("Error while creating raft client session: %v\n", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	req := tuple_spaces.RequestData{Cmd: tuple_spaces.Command_GET, Tuple: &tuple}
	serialized, err := proto.Marshal(&req)

	if err != nil {
		fmt.Printf("Error while serializing request: %v\n", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	for i := 0; i < 10; i += 1 {
		fmt.Printf("Trying to make request: %d/10\n", i+1)
		ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		result, err := nodeHost.SyncPropose(ctx, session, serialized)
		if err == nil {
			foundTuple := tuple_spaces.Tuple{}
			proto.Unmarshal(result.Data, &foundTuple)

			jsonTuple, _ := json.Marshal(foundTuple.Values)
			w.Write(jsonTuple)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Errror: %v\b", err)
			return
		}
		if err != dragonboat.ErrTimeout {
			fmt.Printf("Error while making request: %v", err)
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		}
	}
}

func handleGetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to get all tuples\n")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	session, err := nodeHost.SyncGetSession(ctx, clusterID)
	if err != nil {
		fmt.Printf("Error while creating raft client session: %v\n", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	req := tuple_spaces.RequestData{Cmd: tuple_spaces.Command_READ_ALL}
	serialized, err := proto.Marshal(&req)

	if err != nil {
		fmt.Printf("Error while serializing request: %v\n", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	for i := 0; i < 10; i += 1 {
		fmt.Printf("Trying to make request: %d/10\n", i+1)
		ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		result, err := nodeHost.SyncPropose(ctx, session, serialized)
		if err == nil {
			tupleSpace := tuple_spaces.TupleSpace{}
			proto.Unmarshal(result.Data, &tupleSpace)

			arr := make([][]string, len(tupleSpace.Tuples))
			for i, t := range tupleSpace.Tuples {
				arr[i] = t.Values
			}

			jsonTuple, _ := json.Marshal(&arr)
			w.Write(jsonTuple)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Errror: %v\b", err)
			return
		}
		if err != dragonboat.ErrTimeout {
			fmt.Printf("Error while making request: %v", err)
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		}
	}
}

func createNodeHost(nodeID uint64, initialMembers map[uint64]string) *dragonboat.NodeHost {
	logger.GetLogger("raft").SetLevel(logger.ERROR)
	logger.GetLogger("rsm").SetLevel(logger.ERROR)
	logger.GetLogger("transport").SetLevel(logger.CRITICAL)
	logger.GetLogger("grpc").SetLevel(logger.ERROR)
	logger.GetLogger("dragonboat").SetLevel(logger.ERROR)

	nodeAddr := initialMembers[nodeID]

	cfg := config.Config{
		NodeID:             nodeID,
		ClusterID:          clusterID,
		ElectionRTT:        10,
		HeartbeatRTT:       1,
		CheckQuorum:        true,
		SnapshotEntries:    100,
		CompactionOverhead: 5,
	}

	dataDir := filepath.Join(
		"tuple-space-storage",
		fmt.Sprintf("node-%d", nodeID),
	)

	nodeHostCfg := config.NodeHostConfig{
		WALDir:         dataDir,
		NodeHostDir:    dataDir,
		RTTMillisecond: 200,
		RaftAddress:    nodeAddr,
	}

	nodeHost, err := dragonboat.NewNodeHost(nodeHostCfg)
	if err != nil {
		panic(err)
	}

	if err := nodeHost.StartCluster(initialMembers, false, NewTupleSpaceStateMachine, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start cluster %d: %v", clusterID, err)
		panic(err)
	}
	return nodeHost
}

func main() {
	nodeID := flag.Uint64("nodeid", 1, "Server's ID inside the Raft Cluster. Must be an integer in the interval [1, 5]")
	portRange := flag.Int("port-range", 60000, "Begining of the range of ports allocated by the servers inside the cluster. At least 10 ports must be free inside the interval [port-range, port-range + 19]. E.g: if port-range is 60000, the ports in the interval [60000, 60009] must be free (This argument defaults to 60000)")
	flag.Parse()

	if *nodeID < 1 || *nodeID > 5 {
		fmt.Println("NodeID must be a number between 1 and 5")
	}

	initialMembers := make(map[uint64]string)
	for i := 0; i < 5; i += 1 {
		initialMembers[uint64(i+1)] = fmt.Sprintf("localhost:%d", *portRange+i*2)
	}

	nodeHost = createNodeHost(*nodeID, initialMembers)
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/all", handleGetAll)
	http.HandleFunc("/read", handleRead)
	http.HandleFunc("/", handleHome)

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		printed := false
		tick := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-tick.C:
				leaderID, available, _ := nodeHost.GetLeaderID(clusterID)
				if !available {
					continue
				}

				if leaderID == *nodeID && !printed {
					printed = true
					fmt.Println("Node elected to leader!")
				} else if leaderID != *nodeID && printed {
					printed = false
					fmt.Println("Node demoted to follower!")
				}
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	httpAdress := fmt.Sprintf("localhost:%d", *portRange+int(*nodeID)*2-1)
	fmt.Printf("Server listening at -> %s\n", httpAdress)
	err := http.ListenAndServe(httpAdress, nil)
	cancel()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
