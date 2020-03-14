package memory

import (
	"context"
	"sync"

	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/infix/persistence"
	"github.com/dogmatiq/marshalkit"
)

// Provider is an in-memory implementation of provider.Provider.
type Provider struct {
	dataStores sync.Map
}

// Open returns a data-store for a specific application.
func (p *Provider) Open(
	ctx context.Context,
	app configkit.Identity,
	_ marshalkit.Marshaler,
) (persistence.DataStore, error) {
	ds, _ := p.dataStores.LoadOrStore(
		app.Key,
		&dataStore{},
	)

	return ds.(persistence.DataStore), nil
}