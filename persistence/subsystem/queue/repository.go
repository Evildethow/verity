package queue

import (
	"context"
)

// Repository is an interface for reading persisted event messages.
type Repository interface {
	// LoadQueueMessages loads the next n messages from the queue.
	LoadQueueMessages(ctx context.Context, n int) ([]*Message, error)
}
