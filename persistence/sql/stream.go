package sql

import (
	"context"
	"database/sql"

	"github.com/dogmatiq/infix/envelope"
	"github.com/dogmatiq/infix/persistence"
)

// Stream is a persistence.Stream backed by an SQL database.
type Stream interface {
	persistence.Stream

	// Append atomically appends messages to the stream.
	//
	// It returns the next free offset.
	Append(ctx context.Context, tx *sql.Tx, envelopes ...*envelope.Envelope) (uint64, error)
}
