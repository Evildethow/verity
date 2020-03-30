package eventstore

import "github.com/dogmatiq/infix/draftspecs/envelopespec"

// Offset is the position of an event within the store.
type Offset uint64

// Event is an event persisted in the store.
type Event struct {
	Offset   Offset
	Envelope *envelopespec.Envelope
}