package persistedstream

import (
	"context"

	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/configkit/message"
	"github.com/dogmatiq/infix/eventstream"
	"github.com/dogmatiq/infix/persistence/subsystem/eventstore"
	"github.com/dogmatiq/marshalkit"
)

// Stream is an implementation of Stream that reads events from a
// eventstore.Repository.
type Stream struct {
	// App is the identity of the application that owns the stream.
	App configkit.Identity

	// Types is the set of supported event types.
	Types message.TypeCollection

	// Repository is the event store repository used to query events.
	Repository eventstore.Repository

	// Marshaler is used to unmarshal messages.
	Marshaler marshalkit.Marshaler

	// PreFetch specifies how many messages to pre-load into memory.
	PreFetch int
}

// Application returns the identity of the application that owns the stream.
func (s *Stream) Application() configkit.Identity {
	return s.App
}

// EventTypes returns the set of event types that may appear on the stream.
func (s *Stream) EventTypes(context.Context) (message.TypeCollection, error) {
	return s.Types, nil
}

// Open returns a cursor that reads events from the stream.
//
// o is the offset of the first event to read. The first event on a stream is
// always at offset 0.
//
// f is the set of "filter" event types to be returned by Cursor.Next(). Any
// other event types are ignored.
//
// It returns an error if any of the event types in f are not supported, as
// indicated by EventTypes().
func (s *Stream) Open(
	ctx context.Context,
	o uint64,
	f message.TypeCollection,
) (eventstream.Cursor, error) {
	if f.Len() == 0 {
		panic("at least one event type must be specified")
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	q := eventstore.Query{
		MinOffset: o,
	}

	f.Range(func(mt message.Type) bool {
		n := marshalkit.MustMarshalType(
			s.Marshaler,
			mt.ReflectType(),
		)

		q.Filter.Add(n)

		return true
	})

	consumeCtx, cancelConsume := context.WithCancel(context.Background())

	c := &cursor{
		repository: s.Repository,
		query:      q,
		marshaler:  s.Marshaler,
		cancel:     cancelConsume,
		events:     make(chan *eventstream.Event, s.PreFetch),
	}

	go c.consume(consumeCtx)

	return c, nil
}
