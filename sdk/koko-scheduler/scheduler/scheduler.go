package scheduler

import (
	"context"

	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/coordinator"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/job"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/queue"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/worker"
)

type RegisterUpdate struct {
	GrantedWorkers      []worker.GrantedWorker
	GrantedCoordinators []coordinator.GrantedCoordinator
	Info                []Notice
	Err                 error
}

type Notice struct {
	Code      string
	Message   string
	Requested int
	Granted   int
}

type Config struct {
	Name          string
	ConnectionUrl string
}

type Scheduler struct {
	cfg Config
}

func New(cfg Config) (*Scheduler, error) {
	return &Scheduler{
		cfg: cfg,
	}, nil
}

func (s *Scheduler) registerRequest(ctx context.Context, cfg queue.Config) (<-chan RegisterUpdate, error) {
	return nil, nil
}

func (s *Scheduler) Register(ctx context.Context, cfg queue.Config, handleUpdate func(RegisterUpdate)) (*queue.Queue, error) {
	updates, err := s.registerRequest(ctx, cfg)

	if err != nil {
		return nil, err
	}

	var workers []worker.GrantedWorker
	var coordinators []coordinator.GrantedCoordinator

	for update := range updates {
		if update.Err != nil {
			return nil, update.Err
		}
		workers = append(workers, update.GrantedWorkers...)
		coordinators = append(coordinators, update.GrantedCoordinators...)
		handleUpdate(update)
	}

	return queue.New(cfg, workers, coordinators), nil
}

func (s *Scheduler) Add(ctx context.Context, topic string, payload *job.Job) (*job.Job, error) {
	return nil, nil
}

func (s *Scheduler) Remove(ctx context.Context, topic string, jobID string) error {
	return nil
}

func (s *Scheduler) ViewStatus(ctx context.Context, topic string) {
}

func (s *Scheduler) ViewLogs(ctx context.Context, topic string, jobID string) {
}
