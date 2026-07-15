package scheduler

import "time"

type CoordinatorElectionAlgo string

const (
	CoordinatorElectionAlgoRing   CoordinatorElectionAlgo = "ring"
	CoordinatorElectionAlgoPaxos  CoordinatorElectionAlgo = "paxos"
	CoordinatorElectionAlgoRaft   CoordinatorElectionAlgo = "raft"
	CoordinatorElectionAlgoRandom CoordinatorElectionAlgo = "random"
)

type WorkerConfig struct {
	Count     int
	Segments  []int
	QueueName string
}

type CoordinatorConfig struct {
	Count             int
	Algo              string
	HeartbeatInterval time.Duration
}

type QueueConfig struct {
	Name       string
	RetryDelay time.Duration
}

type Config struct {
	Name          string
	Workers       WorkerConfig
	Coordinators  CoordinatorConfig
	QueueConfig   QueueConfig
	ConnectionUrl string
}

type Scheduler struct {
}

func New(cfg Config) (*Scheduler, error) {
	return &Scheduler{}, nil
}

func (s *Scheduler) Push() {
}
