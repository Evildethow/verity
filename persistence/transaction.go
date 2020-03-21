package persistence

import (
	"context"
	"time"

	"github.com/dogmatiq/dogma"
	"github.com/dogmatiq/infix/envelope"
)

// InstanceRef uniquely identifies an aggregate or process instance at a
// specific revision.
type InstanceRef struct {
	HandlerKey string
	InstanceID string
	Revision   uint64
}

// Transaction exposes persistence operations that can be performed atomically
// in order to handle a single message.
type Transaction interface {
	// PersistAggregate updates (or creates) an aggregate instance.
	PersistAggregate(ctx context.Context, ref InstanceRef, r dogma.AggregateRoot) error

	// PersistProcess updates (or creates) a process instance.
	PersistProcess(ctx context.Context, ref InstanceRef, r dogma.ProcessRoot) error

	// Delete deletes an aggregate or process instance.
	Delete(ctx context.Context, ref InstanceRef) error

	// PersistMessage adds a message to the application's message queue and/or
	// event stream as appropriate.
	PersistMessage(ctx context.Context, env *envelope.Envelope) error

	// Apply applies the changes from the transaction.
	Apply(ctx context.Context) error

	// Abort cancels the transaction, returning the message to the queue.
	//
	// next indicates when the message should be retried.
	Abort(ctx context.Context, next time.Time) error

	// Close closes the transaction.
	Close() error
}
