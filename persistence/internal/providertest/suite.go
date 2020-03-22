package providertest

import (
	"context"
	"time"

	"github.com/dogmatiq/configkit"
	configkitfixtures "github.com/dogmatiq/configkit/fixtures"
	"github.com/dogmatiq/configkit/message"
	"github.com/dogmatiq/dogma"
	dogmafixtures "github.com/dogmatiq/dogma/fixtures"
	"github.com/dogmatiq/infix/persistence"
	"github.com/dogmatiq/marshalkit"
	marshalkitfixtures "github.com/dogmatiq/marshalkit/fixtures"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// In is a container for values that are provided to the provider-specific
// "before" function.
type In struct {
	// Marshaler marshals and unmarshals the test message types, aggregate roots
	// and process roots.
	Marshaler marshalkit.Marshaler
}

// Out is a container for values that are provided by the provider-specific
// "before" function.
type Out struct {
	// Provider is the persistence provider under test.
	Provider persistence.Provider

	// TestTimeout is the maximum duration allowed for each test.
	TestTimeout time.Duration

	// AssumeBlockingDuration specifies how long the tests should wait before
	// assuming a call to Queue.Begin() or Cursor.Next() is successfully
	// blocking, waiting for a new message, as opposed to in the process of
	// "checking" if any messages are already available.
	AssumeBlockingDuration time.Duration
}

const (
	// DefaultTestTimeout is the default test timeout.
	DefaultTestTimeout = 3 * time.Second

	// DefaultAssumeBlockingDuration is the default "assumed blocking duration".
	DefaultAssumeBlockingDuration = 150 * time.Millisecond
)

// Declare declares generic behavioral tests for a specific persistence provider
// implementation.
func Declare(
	before func(context.Context, In) Out,
	after func(),
) {
	var (
		ctx    context.Context
		cancel func()
		in     In
		out    Out

		app1, app2 configkit.RichApplication
	)

	ginkgo.BeforeEach(func() {
		setupCtx, cancelSetup := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelSetup()

		in = In{
			Marshaler: marshalkitfixtures.Marshaler,
		}

		out = before(setupCtx, in)

		if out.TestTimeout <= 0 {
			out.TestTimeout = DefaultTestTimeout
		}

		if out.AssumeBlockingDuration <= 0 {
			out.AssumeBlockingDuration = DefaultAssumeBlockingDuration
		}

		app1 = configkit.FromApplication(&dogmafixtures.Application{
			ConfigureFunc: func(c dogma.ApplicationConfigurer) {
				c.Identity("<app1-name>", "<app1-key>")

				c.RegisterIntegration(&dogmafixtures.IntegrationMessageHandler{
					ConfigureFunc: func(c dogma.IntegrationConfigurer) {
						c.Identity("<int-name>", "<int-key>")

						c.ConsumesCommandType(dogmafixtures.MessageC{})

						c.ProducesEventType(dogmafixtures.MessageE{})
						c.ProducesEventType(dogmafixtures.MessageF{})
						c.ProducesEventType(dogmafixtures.MessageG{})
					},
				})
			},
		})

		app2 = configkit.FromApplication(&dogmafixtures.Application{
			ConfigureFunc: func(c dogma.ApplicationConfigurer) {
				c.Identity("<app2-name>", "<app2-key>")
			},
		})

		ctx, cancel = context.WithTimeout(context.Background(), out.TestTimeout)
	})

	ginkgo.AfterEach(func() {
		if after != nil {
			after()
		}

		cancel()
	})

	ginkgo.Describe("type Provider (interface)", func() {
		ginkgo.Describe("func Open()", func() {
			ginkgo.It("allows repeat calls for the same application", func() {
				ds1, err := out.Provider.Open(ctx, app1, in.Marshaler)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
				defer ds1.Close()

				ds2, err := out.Provider.Open(ctx, app1, in.Marshaler)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
				defer ds2.Close()
			})

			ginkgo.It("returns different instances for different applications", func() {
				ds1, err := out.Provider.Open(ctx, app1, in.Marshaler)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
				defer ds1.Close()

				ds2, err := out.Provider.Open(ctx, app2, in.Marshaler)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
				defer ds2.Close()

				gomega.Expect(ds1).ToNot(gomega.BeIdenticalTo(ds2))
			})
		})
	})

	ginkgo.Describe("type DataStore (interface)", func() {
		var dataStore persistence.DataStore

		ginkgo.BeforeEach(func() {
			var err error
			dataStore, err = out.Provider.Open(ctx, app1, in.Marshaler)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
		})

		ginkgo.AfterEach(func() {
			if dataStore != nil {
				dataStore.Close()
			}
		})

		ginkgo.Describe("func EventStream()", func() {
			ginkgo.It("returns a non-nil event stream", func() {
				s := dataStore.EventStream()
				gomega.Expect(s).ToNot(gomega.BeNil())
			})

			ginkgo.It("returns a stream that advertises the correct types", func() {
				s := dataStore.EventStream()

				types, err := s.MessageTypes(ctx)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())

				gomega.Expect(
					message.IsEqualSetT(
						types,
						message.NewTypeSet(
							configkitfixtures.MessageEType,
							configkitfixtures.MessageFType,
							configkitfixtures.MessageGType,
						),
					),
				).To(gomega.BeTrue())
			})
		})

		ginkgo.Describe("func MessageQueue()", func() {
			ginkgo.It("returns a non-nil message queue", func() {
				q := dataStore.MessageQueue()
				gomega.Expect(q).ToNot(gomega.BeNil())
			})
		})

		ginkgo.Describe("func OffsetRepository()", func() {
			ginkgo.It("returns a non-nil offset repository", func() {
				r := dataStore.OffsetRepository()
				gomega.Expect(r).ToNot(gomega.BeNil())
			})
		})
	})
}
