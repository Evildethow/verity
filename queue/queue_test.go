package queue_test

import (
	"context"
	"errors"
	"time"

	. "github.com/dogmatiq/dogma/fixtures"
	"github.com/dogmatiq/infix/envelope"
	. "github.com/dogmatiq/infix/fixtures"
	. "github.com/dogmatiq/infix/internal/x/gomegax"
	"github.com/dogmatiq/infix/persistence"
	"github.com/dogmatiq/infix/persistence/subsystem/queuestore"
	"github.com/dogmatiq/infix/queue"
	. "github.com/dogmatiq/infix/queue"
	. "github.com/dogmatiq/marshalkit/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// push is a helper function for testing the queue that persists a message to
// the queue then begins tracking it.
func push(
	ctx context.Context,
	q *queue.Queue,
	env *envelope.Envelope,
	nextOptional ...time.Time,
) {
	next := time.Now()
	for _, n := range nextOptional {
		next = n
	}

	p := q.NewParcel(env, next)

	err := persistence.WithTransaction(
		ctx,
		q.DataStore,
		func(tx persistence.ManagedTransaction) error {
			return tx.SaveMessageToQueue(ctx, p)
		},
	)
	Expect(err).ShouldNot(HaveOccurred())

	p.Revision++

	err = q.Track(
		ctx,
		queuestore.Pair{
			Parcel:  p,
			Message: env.Message,
		},
	)
	Expect(err).ShouldNot(HaveOccurred())
}

var _ = Describe("type Queue", func() {
	var (
		ctx              context.Context
		cancel           context.CancelFunc
		dataStore        *DataStoreStub
		repository       *QueueStoreRepositoryStub
		queue            *Queue
		env0, env1, env2 *envelope.Envelope
	)

	BeforeEach(func() {
		ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)

		env0 = NewEnvelope("<message-0>", MessageA1)
		env1 = NewEnvelope("<message-1>", MessageA2)
		env2 = NewEnvelope("<message-2>", MessageA3)

		dataStore = NewDataStoreStub()
		repository = dataStore.QueueStoreRepository().(*QueueStoreRepositoryStub)
		dataStore.QueueStoreRepositoryFunc = func() queuestore.Repository {
			return repository
		}

		queue = &Queue{
			DataStore: dataStore,
			Marshaler: Marshaler,
		}
	})

	AfterEach(func() {
		if dataStore != nil {
			dataStore.Close()
		}

		cancel()
	})

	When("the queue is running", func() {
		JustBeforeEach(func() {
			go queue.Run(ctx)
		})

		Describe("func Pop()", func() {
			When("the queue is empty", func() {
				It("blocks until a message is pushed", func() {
					go func() {
						defer GinkgoRecover()
						time.Sleep(20 * time.Millisecond)
						push(ctx, queue, env0)
					}()

					sess, err := queue.Pop(ctx)
					Expect(err).ShouldNot(HaveOccurred())
					defer sess.Close()
				})

				It("returns an error if the context deadline is exceeded", func() {
					ctx, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
					defer cancel()

					sess, err := queue.Pop(ctx)
					if sess != nil {
						sess.Close()
					}
					Expect(err).To(Equal(context.DeadlineExceeded))
				})
			})

			When("the queue is not empty", func() {
				When("the message at the front of the queue is ready for handling", func() {
					BeforeEach(func() {
						push(ctx, queue, env0)
					})

					It("returns a session immediately", func() {
						ctx, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
						defer cancel()

						sess, err := queue.Pop(ctx)
						Expect(err).ShouldNot(HaveOccurred())
						defer sess.Close()
					})
				})

				When("the message at the front of the queue is not-ready for handling", func() {
					var next time.Time

					BeforeEach(func() {
						next = time.Now().Add(10 * time.Millisecond)
						push(ctx, queue, env0, next)
					})

					It("blocks until the message becomes ready", func() {
						sess, err := queue.Pop(ctx)
						Expect(err).ShouldNot(HaveOccurred())
						defer sess.Close()

						Expect(time.Now()).To(BeTemporally(">=", next))
					})

					It("unblocks if a new message jumps the queue", func() {
						go func() {
							defer GinkgoRecover()
							time.Sleep(5 * time.Millisecond)
							push(ctx, queue, env1)
						}()

						sess, err := queue.Pop(ctx)
						Expect(err).ShouldNot(HaveOccurred())
						defer sess.Close()

						Expect(sess.Envelope()).To(EqualX(
							envelope.MustMarshal(Marshaler, env1),
						))
					})

					It("returns an error if the context deadline is exceeded", func() {
						ctx, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
						defer cancel()

						sess, err := queue.Pop(ctx)
						if sess != nil {
							sess.Close()
						}
						Expect(err).To(Equal(context.DeadlineExceeded))
					})
				})
			})

			When("messages are persisted but not in memory", func() {
				BeforeEach(func() {
					parcels := []*queuestore.Parcel{
						{
							NextAttemptAt: time.Now(),
							Envelope:      envelope.MustMarshal(Marshaler, env0),
						},
						{
							NextAttemptAt: time.Now().Add(10 * time.Millisecond),
							Envelope:      envelope.MustMarshal(Marshaler, env1),
						},
						{
							NextAttemptAt: time.Now().Add(5 * time.Millisecond),
							Envelope:      envelope.MustMarshal(Marshaler, env2),
						},
					}

					err := persistence.WithTransaction(
						ctx,
						dataStore,
						func(tx persistence.ManagedTransaction) error {
							for _, p := range parcels {
								if err := tx.SaveMessageToQueue(ctx, p); err != nil {
									return err
								}
							}
							return nil
						},
					)
					Expect(err).ShouldNot(HaveOccurred())
				})

				It("returns a session for a message loaded from the store", func() {
					sess, err := queue.Pop(ctx)
					Expect(err).ShouldNot(HaveOccurred())
					defer sess.Close()

					Expect(sess.Envelope()).To(EqualX(
						envelope.MustMarshal(Marshaler, env0),
					))
				})
			})
		})

		Describe("func Track()", func() {
			It("panics if the message has not been persisted", func() {
				Expect(func() {
					queue.Track(
						ctx,
						queuestore.Pair{
							Parcel: &queuestore.Parcel{
								Revision: 0,
							},
							Message: env0.Message,
						},
					)
				}).To(Panic())
			})

			It("discards an element if the buffer is full", func() {
				queue.BufferSize = 1

				// This push fills the buffer.
				push(ctx, queue, env0)

				// This push exceeds the limit so env1 should not be buffered.
				push(ctx, queue, env1)

				// Acquire a session for env0, but don't commit it.
				sess, err := queue.Pop(ctx)
				Expect(err).ShouldNot(HaveOccurred())
				defer sess.Close()

				// Nothing new will be loaded from the store while there is
				// anything tracked at all (this is why its important to
				// configure the buffer size larger than the number of
				// consumers).
				ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
				defer cancel()

				sess, err = queue.Pop(ctx)
				if sess != nil {
					sess.Close()
				}
				Expect(err).To(Equal(context.DeadlineExceeded))
			})

			When("a message is tracked while loading from the store", func() {
				It("does not duplicate the message", func() {
					repository.LoadQueueMessagesFunc = func(
						ctx context.Context,
						n int,
					) ([]*queuestore.Parcel, error) {
						push(ctx, queue, env0)
						return repository.Repository.LoadQueueMessages(ctx, n)
					}

					// We expect to get the pushed message once.
					sess, err := queue.Pop(ctx)
					Expect(err).ShouldNot(HaveOccurred())
					defer sess.Close()

					ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
					defer cancel()

					// But not twice.
					sess, err = queue.Pop(ctx)
					if sess != nil {
						sess.Close()
					}
					Expect(err).To(Equal(context.DeadlineExceeded))
				})
			})
		})
	})

	When("the queue is not running", func() {
		Describe("func Track()", func() {
			It("returns an error if the deadline is exceeded", func() {
				p := queuestore.Pair{
					Parcel: &queuestore.Parcel{
						Revision: 1,
						Envelope: envelope.MustMarshal(Marshaler, env0),
					},
					Message: env0.Message,
				}

				// It's an implementation detail, but the internal channel used to start
				// tracking is buffered at the same size as the overall buffer size
				// limit.
				//
				// We can't set it to zero, because that will fallback to the default.
				// We also can't start the queue, otherwise it'll start reading from
				// this channel and nothing will block.
				//
				// Instead, we set it to one, and "fill" the channel with a request to
				// ensure that it will block.
				queue.BufferSize = 1
				err := queue.Track(ctx, p)
				Expect(err).ShouldNot(HaveOccurred())

				// Setup a short deadline for the test.
				ctx, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
				defer cancel()

				err = queue.Track(ctx, p)
				Expect(err).To(Equal(context.DeadlineExceeded))
			})
		})

		Describe("fun Run()", func() {
			It("returns an error if messages can not be loaded from the repository", func() {
				repository.LoadQueueMessagesFunc = func(
					context.Context,
					int,
				) ([]*queuestore.Parcel, error) {
					return nil, errors.New("<error>")
				}

				err := queue.Run(ctx)
				Expect(err).To(MatchError("<error>"))
			})
		})
	})

	When("the queue has stopped", func() {
		BeforeEach(func() {
			ctx, cancel := context.WithCancel(context.Background())
			cancel() // cancel immediately

			queue.Run(ctx)
		})

		Describe("func Track()", func() {
			It("does not block", func() {
				p := queuestore.Pair{
					Parcel: &queuestore.Parcel{
						Revision: 1,
						Envelope: envelope.MustMarshal(Marshaler, env0),
					},
					Message: env0.Message,
				}

				err := queue.Track(ctx, p)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
