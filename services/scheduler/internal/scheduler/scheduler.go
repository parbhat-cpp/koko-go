package scheduler

import (
	"context"
	"log"

	pb "github.com/parbhat-cpp/koko-go/proto/gen/scheduler/v1"
	shared "github.com/parbhat-cpp/koko-go/proto/gen/shared/v1"
	rdb "github.com/parbhat-cpp/koko-go/services/scheduler/internal/redis"
	"github.com/parbhat-cpp/koko-go/services/scheduler/internal/utils"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SchedulerServiceServer struct {
	pb.UnimplementedSchedulerServiceServer
	scheduler *Scheduler
}

type Scheduler struct {
	Rdb  *redis.Client
	Name string
}

func New(rdb *redis.Client, name string) *Scheduler {
	return &Scheduler{
		Rdb:  rdb,
		Name: name,
	}
}

func (s *SchedulerServiceServer) Init(ctx context.Context, req *pb.SchedulerServiceInitRequest) (*pb.SchedulerServiceInitResponse, error) {
	log.Println("start")
	log.Printf("Request: %v %v", req.Name, req.RedisConfig)
	if s.scheduler != nil {
		log.Printf("Scheduler already initialized: %v", s.scheduler.Name)
		return &pb.SchedulerServiceInitResponse{
			Success: false,
			Name:    &shared.ServiceName{Name: ""},
		}, status.Error(codes.AlreadyExists, "Scheduler already initialized")
	}

	rdbCfg := utils.ToRedisOptions(req.RedisConfig)
	r, err := rdb.Connect(*rdbCfg)

	if err != nil {
		return nil, err
	}

	scheduler := New(r, req.Name.Name)

	s.scheduler = scheduler

	log.Printf("done")

	return &pb.SchedulerServiceInitResponse{
		Success: true,
		Name:    &shared.ServiceName{Name: scheduler.Name},
	}, nil
}

func (s *SchedulerServiceServer) Register(req *pb.RegisterRequest, stream pb.SchedulerService_RegisterServer) error {
	return nil
}

func (s *SchedulerServiceServer) AddJob(ctx context.Context, req *pb.AddJobRequest) (*pb.AddJobResponse, error) {
	return nil, status.Error(codes.Unimplemented, "AddJob method not implemented yet")
}

func (s *SchedulerServiceServer) RemoveJob(ctx context.Context, req *pb.RemoveJobRequest) (*pb.RemoveJobResponse, error) {
	return nil, status.Error(codes.Unimplemented, "RemoveJob method not implemented yet")
}

func (s *SchedulerServiceServer) ViewStatus(ctx context.Context, req *pb.ViewStatusRequest) (*pb.ViewStatusResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ViewStatus method not implemented yet")
}
