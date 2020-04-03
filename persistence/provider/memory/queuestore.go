package memory

import (
	"context"
	"errors"
	"sort"

	"github.com/dogmatiq/infix/persistence/subsystem/queuestore"
)

// SaveMessageToQueue persists a message to the application's message queue.
//
// If the message is already on the queue its meta-data is updated.
//
// m.Revision must be the revision of the message as currently persisted,
// otherwise an optimistic concurrency conflict has occurred, the message
// is not saved and ErrConflict is returned.
func (t *transaction) SaveMessageToQueue(
	ctx context.Context,
	m *queuestore.Message,
) (err error) {
	if err := t.begin(ctx); err != nil {
		return err
	}

	id := m.Envelope.GetMetaData().GetMessageId()

	var rev queuestore.Revision
	if x, ok := t.uncommitted.queue[id]; ok {
		rev = x.Revision
	} else if x, ok := t.ds.db.queue.uniq[id]; ok {
		rev = x.Revision
	}

	if m.Revision != rev {
		return queuestore.ErrConflict
	}

	m = cloneQueueMessage(m)
	m.Revision++

	if t.uncommitted.queue == nil {
		t.uncommitted.queue = map[string]*queuestore.Message{}
	}
	t.uncommitted.queue[id] = m

	return nil
}

// commitQueue commits staged queue items to the database.
func (t *transaction) commitQueue() {
	q := &t.ds.db.queue

	for id, m := range t.uncommitted.queue {
		if q.uniq == nil {
			q.uniq = map[string]*queuestore.Message{}
		} else if x, ok := q.uniq[id]; ok {
			// This message is already on the queue, we don't insert it, only
			// update its meta-data.
			x.Revision = m.Revision

			if !x.NextAttemptAt.Equal(m.NextAttemptAt) {
				x.NextAttemptAt = m.NextAttemptAt
				sort.Slice(
					q.order,
					func(i, j int) bool {
						return q.order[i].NextAttemptAt.Before(
							q.order[j].NextAttemptAt,
						)
					},
				)
			}

			continue
		}

		// Add the message to the unique index.
		q.uniq[id] = m

		// Find the index where we'll insert our event. It's the index of the
		// first message that has a NextAttemptedAt greater than m's.
		index := sort.Search(
			len(q.order),
			func(i int) bool {
				return m.NextAttemptAt.Before(
					q.order[i].NextAttemptAt,
				)
			},
		)

		// Expand the size of the queue.
		q.order = append(q.order, nil)

		// Shift messages further back to make space for m.
		copy(q.order[index+1:], q.order[index:])

		// Insert m at the index.
		q.order[index] = m
	}
}

// RemoveMessageFromQueue removes a specific message from the application's
// message queue.
//
// m.Revision must be the revision of the message as currently persisted,
// otherwise an optimistic concurrency conflict has occurred, the message
// remains on the queue and ErrConflict is returned.
func (t *transaction) RemoveMessageFromQueue(
	ctx context.Context,
	m *queuestore.Message,
) (err error) {
	return errors.New("not implemented")
}

// queueStoreRepository is an implementation of queuestore.Repository that
// stores queued messages in memory.
type queueStoreRepository struct {
	db *database
}

// LoadQueueMessages loads the next n messages from the queue.
func (r *queueStoreRepository) LoadQueueMessages(
	ctx context.Context,
	n int,
) ([]*queuestore.Message, error) {
	if err := r.db.RLock(ctx); err != nil {
		return nil, err
	}
	defer r.db.RUnlock()

	max := len(r.db.queue.order)
	if n > max {
		n = max
	}

	result := make([]*queuestore.Message, n)

	for i, m := range r.db.queue.order[:n] {
		result[i] = cloneQueueMessage(m)
	}

	return result, nil
}
