package scheduler

import (
	"context"
	"fmt"

	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/coordinator"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/job"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/queue"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/utils"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/worker"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/parbhat-cpp/koko-go/proto/gen/scheduler/v1"
	sharedv1 "github.com/parbhat-cpp/koko-go/proto/gen/shared/v1"
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
	Name        string
	RedisConfig redis.Options
	ConnUrl     string
}

type Scheduler struct {
	cfg    Config
	conn   *grpc.ClientConn
	client pb.SchedulerServiceClient
}

func New(cfg Config) (*Scheduler, error) {
	conn, err := grpc.NewClient(cfg.ConnUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	client := pb.NewSchedulerServiceClient(conn)

	res, err := client.Init(context.Background(), &pb.SchedulerServiceInitRequest{
		Name:        &sharedv1.ServiceName{Name: cfg.Name},
		RedisConfig: utils.ToSchedulerRedisOptions(&cfg.RedisConfig),
	})

	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	if !res.Success {
		_ = conn.Close()
		return nil, fmt.Errorf("scheduler init failed for %s", cfg.Name)
	}

	return &Scheduler{
		cfg:    cfg,
		conn:   conn,
		client: client,
	}, nil
}

func (s *Scheduler) Close() error {
	if s.conn == nil {
		return nil
	}

	return s.conn.Close()
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
