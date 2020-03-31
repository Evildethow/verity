package memory

import (
	"github.com/dogmatiq/infix/draftspecs/envelopespec"
	"github.com/dogmatiq/infix/persistence/subsystem/eventstore"
	"github.com/dogmatiq/infix/persistence/subsystem/queuestore"
	"github.com/golang/protobuf/proto"
)

func cloneEvent(ev *eventstore.Event) *eventstore.Event {
	clone := *ev
	clone.Envelope = cloneEnvelope(clone.Envelope)
	return &clone
}

func cloneQueueMessage(m *queuestore.Message) *queuestore.Message {
	clone := *m
	clone.Envelope = cloneEnvelope(clone.Envelope)
	return &clone
}

func cloneEnvelope(env *envelopespec.Envelope) *envelopespec.Envelope {
	return proto.Clone(env).(*envelopespec.Envelope)
}
