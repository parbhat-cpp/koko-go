package coordinator

import "time"

type GrantedCoordinator struct {
	ID        string
	Role      string
	QueueName string
	CreatedAt time.Time
}
