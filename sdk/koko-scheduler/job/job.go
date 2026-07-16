package job

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	id         string
	MaxRetries int
	Payload    any
	RunAt      time.Time
}

func (j *Job) ID() string {
	return j.id
}

func New(payload any, runAt time.Time, maxRetries int) *Job {
	return &Job{
		id:         uuid.New().String(),
		MaxRetries: maxRetries,
		Payload:    payload,
		RunAt:      runAt,
	}
}
