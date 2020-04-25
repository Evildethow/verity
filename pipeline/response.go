package pipeline

import (
	"context"
	"time"

	"github.com/dogmatiq/infix/internal/pooling"
	"github.com/dogmatiq/infix/parcel"
	"github.com/dogmatiq/infix/persistence"
	"github.com/dogmatiq/infix/persistence/subsystem/eventstore"
	"github.com/dogmatiq/infix/persistence/subsystem/queuestore"
)

// Response is the result from a pipeline stage.
//
// It encapsulates the messages that were produced, so they may be observed by
// other components of the engine.
type Response struct {
	queueParcels []*parcel.Parcel
	queueItems   []*queuestore.Item
	eventParcels []*parcel.Parcel
	eventItems   []*eventstore.Item
}

// EnqueueMessage is a helper method that adds a message to the queue and
// adds it to the response.
func (r *Response) EnqueueMessage(
	ctx context.Context,
	tx persistence.ManagedTransaction,
	p *parcel.Parcel,
) error {
	n := p.ScheduledFor
	if n.IsZero() {
		n = time.Now()
	}

	i := &queuestore.Item{
		NextAttemptAt: n,
		Envelope:      p.Envelope,
	}

	if err := tx.SaveMessageToQueue(ctx, i); err != nil {
		return err
	}

	i.Revision++

	if r.queueParcels == nil {
		r.queueParcels = pooling.ParcelSlice.Get(1)
		r.queueItems = pooling.QueueStoreItemSlice.Get(1)
	}

	r.queueParcels = append(r.queueParcels, p)
	r.queueItems = append(r.queueItems, i)

	return nil
}

// RecordEvent is a helper method that appends an event to the event stream and
// adds it to the response.
func (r *Response) RecordEvent(
	ctx context.Context,
	tx persistence.ManagedTransaction,
	p *parcel.Parcel,
) (uint64, error) {
	o, err := tx.SaveEvent(ctx, p.Envelope)
	if err != nil {
		return 0, err
	}

	i := &eventstore.Item{
		Offset:   o,
		Envelope: p.Envelope,
	}

	if r.eventParcels == nil {
		r.eventParcels = pooling.ParcelSlice.Get(1)
		r.eventItems = pooling.EventStoreItemSlice.Get(1)
	}

	r.eventParcels = append(r.eventParcels, p)
	r.eventItems = append(r.eventItems, i)

	return o, nil
}
