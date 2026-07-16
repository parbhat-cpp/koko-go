package main

import (
	"context"
	"time"

	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/job"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/queue"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/scheduler"
	"github.com/parbhat-cpp/koko-go/sdk/koko-scheduler/worker"
)

func main() {
	s, err := scheduler.New(scheduler.Config{
		Name:          "Koko",
		ConnectionUrl: "redis://localhost:6379",
	})

	q, err := s.Register(context.Background(), queue.Config{
		Topic: "email:service",
		WorkerConfig: &queue.WorkerConfig{
			Count: 3,
		},
		CoordinatorConfig: &queue.CoordinatorConfig{
			Count:             2,
			Algo:              queue.CoordinatorElectionAlgoRaft,
			HeartbeatInterval: time.Minute,
		},
	}, nil)

	if err != nil {
		panic(err)
	}

	j := job.New(map[string]any{
		"name":     "John Doe",
		"template": "welcome-email",
	}, time.Now().Add(time.Minute*5), 3)

	s.Add(context.Background(), "email:service", j)

	s.Add(context.Background(), "email:service", &job.Job{
		MaxRetries: 3,
		Payload: map[string]any{
			"name":     "John Doe",
			"template": "welcome-email",
		},
		RunAt: time.Now().Add(time.Minute * 5),
	})

	q.Worker(context.Background(), worker.Worker{
		Handler: func(i interface{}) error {
			// do something
			return nil
		},
	})
}
