package streamtest

import (
	"context"
	"sync"
	"time"

	configkitfixtures "github.com/dogmatiq/configkit/fixtures"
	"github.com/dogmatiq/configkit/message"
	dogmafixtures "github.com/dogmatiq/dogma/fixtures"
	"github.com/dogmatiq/infix/envelope"
	infixfixtures "github.com/dogmatiq/infix/fixtures"
	"github.com/dogmatiq/infix/persistence"
	"github.com/dogmatiq/linger"
	"github.com/dogmatiq/marshalkit"
	marshalkitfixtures "github.com/dogmatiq/marshalkit/fixtures"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// Config is the result of setting up a test suite.
type Config struct {
	// Stream is the stream to be tested.
	Stream persistence.Stream

	// TestTimeout is the maximum duration allowed for each test.
	TestTimeout time.Duration

	// AssumeBlockingDuration specifies how long the tests should wait before
	// assuming a call to Cursor.Next() is successfully blocking, waiting for a
	// new message, as opposed to in the process of "checking" if any messages
	// are already available.
	AssumeBlockingDuration time.Duration

	// Append is a function that appends messages to the stream.
	Append func(context.Context, ...*envelope.Envelope)
}

const (
	// DefaultTestTimeout is the default test timeout.
	DefaultTestTimeout = 3 * time.Second

	// DefaultAssumeBlockingDuration is the default "assumed blocking duration".
	DefaultAssumeBlockingDuration = 150 * time.Millisecond
)

// Declare declares generic behavioral tests for a specific driver
// implementation.
func Declare(
	before func(context.Context, marshalkit.Marshaler) Config,
	after func(),
) {
	var (
		ctx    context.Context
		cancel func()
		cfg    Config
		types  message.TypeSet

		env0 = infixfixtures.NewEnvelope("<message-0>", dogmafixtures.MessageA1)
		env1 = infixfixtures.NewEnvelope("<message-1>", dogmafixtures.MessageB1)
		env2 = infixfixtures.NewEnvelope("<message-2>", dogmafixtures.MessageA2)
		env3 = infixfixtures.NewEnvelope("<message-3>", dogmafixtures.MessageB2)
		env4 = infixfixtures.NewEnvelope("<message-4>", dogmafixtures.MessageC1)

		message0 = &persistence.StreamMessage{Offset: 0, Envelope: env0}
		message1 = &persistence.StreamMessage{Offset: 1, Envelope: env1}
		message2 = &persistence.StreamMessage{Offset: 2, Envelope: env2}
		message3 = &persistence.StreamMessage{Offset: 3, Envelope: env3}
		message4 = &persistence.StreamMessage{Offset: 4, Envelope: env4}
	)

	ginkgo.BeforeEach(func() {
		setupCtx, cancelSetup := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelSetup()
		cfg = before(setupCtx, marshalkitfixtures.Marshaler)

		if cfg.TestTimeout <= 0 {
			cfg.TestTimeout = DefaultTestTimeout
		}

		if cfg.AssumeBlockingDuration <= 0 {
			cfg.AssumeBlockingDuration = DefaultAssumeBlockingDuration
		}

		ctx, cancel = context.WithTimeout(context.Background(), cfg.TestTimeout)

		types = message.NewTypeSet(
			configkitfixtures.MessageAType,
			configkitfixtures.MessageBType,
			configkitfixtures.MessageCType,
		)
	})

	ginkgo.AfterEach(func() {
		if after != nil {
			after()
		}

		cancel()
	})

	ginkgo.Describe("type Stream", func() {
		ginkgo.Describe("func Open()", func() {
			ginkgo.BeforeEach(func() {
				cfg.Append(
					ctx,
					env0,
					env1,
					env2,
					env3,
				)
			})

			ginkgo.It("honours the initial offset", func() {
				cur, err := cfg.Stream.Open(ctx, 2, types)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
				defer cur.Close()

				m, err := cur.Next(ctx)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
				gomega.Expect(m).To(gomega.Equal(message2))
			})

			ginkgo.It("limits results to the supplied message types", func() {
				types = message.NewTypeSet(
					configkitfixtures.MessageAType,
				)

				cur, err := cfg.Stream.Open(ctx, 0, types)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
				defer cur.Close()

				m, err := cur.Next(ctx)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
				gomega.Expect(m).To(gomega.Equal(message0))

				m, err = cur.Next(ctx)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
				gomega.Expect(m).To(gomega.Equal(message2))
			})

			ginkgo.It("returns an error if the context is canceled", func() {
				cancel()

				_, err := cfg.Stream.Open(ctx, 0, types)
				gomega.Expect(err).Should(gomega.HaveOccurred())
			})
		})
	})

	ginkgo.Describe("type Cursor", func() {
		ginkgo.Describe("func Next()", func() {
			ginkgo.Context("when the stream is empty", func() {
				ginkgo.It("blocks", func() {
					cur, err := cfg.Stream.Open(ctx, 0, types)
					gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
					defer cur.Close()

					ctx, cancel := context.WithTimeout(ctx, cfg.AssumeBlockingDuration)
					defer cancel()

					_, err = cur.Next(ctx)
					gomega.Expect(err).To(gomega.Equal(context.DeadlineExceeded))
				})
			})

			ginkgo.Context("when the stream is not empty", func() {
				ginkgo.BeforeEach(func() {
					cfg.Append(
						ctx,
						env0,
						env1,
						env2,
						env3,
					)
				})

				ginkgo.It("returns the messages in order", func() {
					cur, err := cfg.Stream.Open(ctx, 0, types)
					gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
					defer cur.Close()

					m, err := cur.Next(ctx)
					gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
					gomega.Expect(m).To(gomega.Equal(message0))

					m, err = cur.Next(ctx)
					gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
					gomega.Expect(m).To(gomega.Equal(message1))

					m, err = cur.Next(ctx)
					gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
					gomega.Expect(m).To(gomega.Equal(message2))

					m, err = cur.Next(ctx)
					gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
					gomega.Expect(m).To(gomega.Equal(message3))
				})

				ginkgo.It("returns an error if the cursor is closed", func() {
					cur, err := cfg.Stream.Open(ctx, 0, types)
					gomega.Expect(err).ShouldNot(gomega.HaveOccurred())

					cur.Close()

					_, err = cur.Next(ctx)
					gomega.Expect(err).To(gomega.Equal(persistence.ErrStreamCursorClosed))
				})

				ginkgo.It("returns an error if the context is canceled", func() {
					cur, err := cfg.Stream.Open(ctx, 4, types)
					gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
					defer cur.Close()

					cancel()

					_, err = cur.Next(ctx)
					gomega.Expect(err).Should(gomega.HaveOccurred())
				})

				ginkgo.When("waiting for a new message", func() {
					ginkgo.It("wakes if a message is appended", func() {
						// Open a cursor after the offset of the existing messages.
						cur, err := cfg.Stream.Open(ctx, 4, types)
						gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
						defer cur.Close()

						go func() {
							time.Sleep(cfg.AssumeBlockingDuration)
							cfg.Append(ctx, env4)
						}()

						m, err := cur.Next(ctx)
						gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
						gomega.Expect(m).To(gomega.Equal(message4))
					})

					ginkgo.It("returns an error if the cursor is closed", func() {
						cur, err := cfg.Stream.Open(ctx, 4, types)
						gomega.Expect(err).ShouldNot(gomega.HaveOccurred())

						go func() {
							time.Sleep(cfg.AssumeBlockingDuration)
							cur.Close()
						}()

						_, err = cur.Next(ctx)
						gomega.Expect(err).To(gomega.Equal(persistence.ErrStreamCursorClosed))
					})

					ginkgo.It("returns an error if the context is canceled", func() {
						cur, err := cfg.Stream.Open(ctx, 4, types)
						gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
						defer cur.Close()

						go func() {
							time.Sleep(cfg.AssumeBlockingDuration)
							cancel()
						}()

						_, err = cur.Next(ctx)
						gomega.Expect(err).Should(gomega.HaveOccurred())
					})

					ginkgo.It("does not compete with other waiting cursors", func() {
						// This test ensures that when there are multiple
						// cursors awaiting a new message they are all woken
						// when a message is appended.

						const cursors = 3

						// barrier is used to delay the append until all of the
						// cursors have started blocking.
						barrier := make(chan struct{}, cursors)

						var g sync.WaitGroup
						defer g.Wait()

						g.Add(cursors)

						// start the cursors
						for i := 0; i < cursors; i++ {
							go func() error {
								linger.SleepX(
									ctx,
									linger.FullJitter,
									cfg.AssumeBlockingDuration,
								)

								defer g.Done()
								defer ginkgo.GinkgoRecover()

								cur, err := cfg.Stream.Open(ctx, 4, types)
								gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
								defer cur.Close()

								barrier <- struct{}{}
								m, err := cur.Next(ctx)
								gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
								gomega.Expect(m).To(gomega.Equal(message4))

								return nil
							}()
						}

						// wait for the cursors to signal they are about to block
						for i := 0; i < cursors; i++ {
							select {
							case <-barrier:
							case <-ctx.Done():
								gomega.Expect(ctx.Err()).ShouldNot(gomega.HaveOccurred())
							}
						}

						time.Sleep(cfg.AssumeBlockingDuration)

						// wake the consumers
						cfg.Append(ctx, env4)
					})
				})
			})
		})

		ginkgo.Describe("func Close()", func() {
			ginkgo.It("returns an error if the cursor is already closed", func() {
				cur, err := cfg.Stream.Open(ctx, 4, types)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())

				err = cur.Close()
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())

				err = cur.Close()
				gomega.Expect(err).To(gomega.Equal(persistence.ErrStreamCursorClosed))
			})
		})
	})
}