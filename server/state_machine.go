package main

import (
	"errors"
	"fmt"
	"io"

	tuple_spaces "computacao-distribuida/tuple-spaces"

	"github.com/lni/dragonboat/v3/logger"
	"github.com/lni/dragonboat/v3/statemachine"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var logging = logger.GetLogger("computacao_distribuida")

// Vê se as tuplas `a` e `pattern` são equivalentes
func match(a *tuple_spaces.Tuple, pattern *tuple_spaces.Tuple) bool {
	if len(a.GetValues()) != len(pattern.GetValues()) {
		return false
	}

	for i, val := range pattern.GetValues() {
		if val != a.GetValues()[i] && val != "*" {
			return false
		}
	}
	return true
}

// retorna uma tupla rquivalente à `tuple_pattern` se ela existir. Senão, retorna `nil`
func get(ts *tuple_spaces.TupleSpace, tuple_pattern *tuple_spaces.Tuple, remove bool) *tuple_spaces.Tuple {
	for i, t := range ts.GetTuples() {
		if len(t.GetValues()) != len(tuple_pattern.GetValues()) {
			continue
		} else if match(t, tuple_pattern) {
			if remove {
				// Magic WoW -> in C, that would be a seg fault
				tuple := t
				ts.Tuples = append(ts.Tuples[:i], ts.Tuples[i+1:]...)
				return tuple
			} else {
				return t
			}
		}
	}
	return nil
}

var (
	SerializationError   = errors.New("Error while serializing tuple")
	DeserializationError = errors.New("Error while deserializing tuple")
)

type TupleSpaceStateMachine struct {
	tupleSpace tuple_spaces.TupleSpace
	clusterID  uint64
	nodeID     uint64
}

func NewTupleSpaceStateMachine(clusterID uint64, nodeID uint64) statemachine.IStateMachine {
	return &TupleSpaceStateMachine{clusterID: clusterID, nodeID: nodeID, tupleSpace: tuple_spaces.TupleSpace{}}
}

// LookUp não muda o estado. No nosso caso, equivale ao Read(tuple). I.e: bustca uma tupla, mas não a remove
func (s *TupleSpaceStateMachine) Lookup(query interface{}) (interface{}, error) {
	if query == nil {
		logging.Warningf("Nil value received in LookUp")
		return nil, nil
	}

	tuple, ok := query.(*tuple_spaces.Tuple)
	if !ok {
		logging.Errorf("Error: value in LookUp is not *Tuple!")
		return nil, DeserializationError
	}
	logging.Infof("Reading tuple: %v", tuple)

	if result := get(&s.tupleSpace, tuple, false); result != nil {
		return result, nil
	}
	fmt.Print("Pattern not found")
	return nil, nil
}

func (s *TupleSpaceStateMachine) Update(data []byte) (statemachine.Result, error) {
	if data == nil {
		logging.Warningf("Nil value received in Update")
		return statemachine.Result{
			Value: 1, // Failure
		}, nil
	}

	request := &tuple_spaces.RequestData{}
	if err := proto.Unmarshal(data, request); err != nil {
		logging.Errorf("Error while deserializing request: %v", err)
		return statemachine.Result{Value: 1}, nil
	}

	var result protoreflect.ProtoMessage
	switch request.GetCmd() {
	case tuple_spaces.Command_GET:
		logging.Infof("Command received: GET")
		result = get(&s.tupleSpace, request.GetTuple(), true)
	case tuple_spaces.Command_READ:
		logging.Infof("Command received: READ")
		result = get(&s.tupleSpace, request.GetTuple(), false)
	case tuple_spaces.Command_WRITE:
		logging.Infof("Command received: WRITE")
		s.tupleSpace.Tuples = append(s.tupleSpace.Tuples, request.Tuple)
		return statemachine.Result{Value: 0}, nil
	case tuple_spaces.Command_READ_ALL:
		logging.Infof("Command received: READ ALL")
		result = &s.tupleSpace
	default:
		logging.Errorf("Error, invalid command received: %d", request.Cmd)
		return statemachine.Result{Value: 2}, nil
	}

	if result == nil {
		return statemachine.Result{
			Value: 0, // success
		}, nil
	}

	serializedResult, err := proto.Marshal(result)
	if err != nil {
		logging.Errorf("Error while serializing tuple: %v", err)
		return statemachine.Result{Value: 1}, nil
	}
	return statemachine.Result{Value: 0, Data: serializedResult}, nil
}

func (s *TupleSpaceStateMachine) SaveSnapshot(w io.Writer, fc statemachine.ISnapshotFileCollection, done <-chan struct{}) error {
	serialized, err := proto.Marshal(&s.tupleSpace)
	if err != nil {
		logging.Errorf("Error while serializing tuple space: %v", err)
		return nil
	}
	_, err = w.Write(serialized)
	return err
}

func (s *TupleSpaceStateMachine) RecoverFromSnapshot(r io.Reader, files []statemachine.SnapshotFile, done <-chan struct{}) error {
	data, err := io.ReadAll(r)
	if err != nil {
		logging.Errorf("Error while recovering from snapshot: %v", err)
		return err
	}

	s.tupleSpace = tuple_spaces.TupleSpace{}
	if err := proto.Unmarshal(data, &s.tupleSpace); err != nil {
		logging.Errorf("Error while deserializing tuple space: %v", err)
	}
	return nil
}

func (s *TupleSpaceStateMachine) Close() error { return nil }
