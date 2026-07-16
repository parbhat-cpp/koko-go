package worker

import "time"

type Worker struct {
	Handler func(interface{}) error
}

type GrantedWorker struct {
	ID        string
	Segments  []int
	QueueName string
	CreatedAt time.Time
}
