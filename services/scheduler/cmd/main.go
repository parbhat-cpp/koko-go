package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/parbhat-cpp/koko-go/proto/gen/scheduler/v1"
	"github.com/parbhat-cpp/koko-go/services/scheduler/internal/scheduler"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Server running on port: 50051")

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSchedulerServiceServer(grpcServer, &scheduler.SchedulerServiceServer{})
	grpcServer.Serve(lis)
}
