package queue

import (
	"context"
	"sync"
	"time"

	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/coordinator"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/worker"
)

type CoordinatorElectionAlgo string

const (
	CoordinatorElectionAlgoRing   CoordinatorElectionAlgo = "ring"
	CoordinatorElectionAlgoPaxos  CoordinatorElectionAlgo = "paxos"
	CoordinatorElectionAlgoRaft   CoordinatorElectionAlgo = "raft"
	CoordinatorElectionAlgoRandom CoordinatorElectionAlgo = "random"
)

type WorkerConfig struct {
	Count     int
	segments  []int
	queueName string
}

type CoordinatorConfig struct {
	Count             int
	Algo              CoordinatorElectionAlgo
	HeartbeatInterval time.Duration
}

type Config struct {
	Topic             string
	WorkerConfig      *WorkerConfig
	CoordinatorConfig *CoordinatorConfig
}

type Queue struct {
	mu           sync.RWMutex
	cfg          Config
	workers      []worker.GrantedWorker
	coordinators []coordinator.GrantedCoordinator
}

func New(cfg Config, workers []worker.GrantedWorker, coordinators []coordinator.GrantedCoordinator) *Queue {
	return &Queue{
		cfg:          cfg,
		workers:      workers,
		coordinators: coordinators,
	}
}

func (q *Queue) Worker(ctx context.Context, wkr worker.Worker) (*worker.Worker, error) {
	return nil, nil
}
