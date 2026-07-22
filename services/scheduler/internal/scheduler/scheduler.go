package scheduler

import (
	"context"

	pb "github.com/parbhat-cpp/koko-go/proto/gen/scheduler/v1"
	"github.com/parbhat-cpp/koko-go/services/scheduler/internal/redis"
	"github.com/parbhat-cpp/koko-go/services/scheduler/internal/utils"
	go_redis "github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SchedulerServiceServer struct {
	pb.UnimplementedSchedulerServiceServer
	scheduler *Scheduler
}

type Scheduler struct {
	Rdb  *go_redis.Client
	Name string
}

func New(rdb *go_redis.Client, name string) *Scheduler {
	return &Scheduler{
		Rdb:  rdb,
		Name: name,
	}
}

func (s *SchedulerServiceServer) Init(ctx context.Context, req *pb.SchedulerServiceInitRequest) (*pb.SchedulerServiceInitResponse, error) {
	if s.scheduler != nil {
		return &pb.SchedulerServiceInitResponse{
			Success: false,
			Name:    "",
		}, status.Error(codes.AlreadyExists, "Scheduler already initialized")
	}

	rdbCfg := utils.ToRedisOptions(req.RedisConfig)
	rdb, err := redis.Connect(*rdbCfg)

	if err != nil {
		return nil, err
	}

	scheduler := New(rdb, req.Name)

	s.scheduler = scheduler

	return &pb.SchedulerServiceInitResponse{
		Success: true,
		Name:    scheduler.Name,
	}, nil
}
