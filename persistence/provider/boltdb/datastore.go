package boltdb

import (
	"context"
	"sync"

	"github.com/dogmatiq/infix/internal/x/bboltx"
	"github.com/dogmatiq/infix/persistence"
	"github.com/dogmatiq/infix/persistence/eventstore"
)

// dataStore is an implementation of persistence.DataStore for BoltDB.
type dataStore struct {
	db     *database
	appKey []byte

	m       sync.Mutex
	txns    int
	closed  bool
	done    chan struct{}
	release func(string) error
}

func newDataStore(
	db *database,
	k string,
	r func(string) error,
) *dataStore {
	return &dataStore{
		db:      db,
		appKey:  []byte(k),
		done:    make(chan struct{}),
		release: r,
	}
}

// EventStoreRepository returns the application's event store repository.
func (ds *dataStore) EventStoreRepository() eventstore.Repository {
	return &eventStoreRepository{ds.db, ds.appKey}
}

// Begin starts a new transaction.
func (ds *dataStore) Begin(ctx context.Context) (_ persistence.Transaction, err error) {
	defer bboltx.Recover(&err)

	ds.m.Lock()
	defer ds.m.Unlock()

	if ds.closed {
		return nil, persistence.ErrDataStoreClosed
	}

	ds.txns++

	return &transaction{
		db:      ds.db,
		appKey:  ds.appKey,
		release: ds.releaseTx,
	}, nil
}

// Close closes the data store.
//
// Closing a data-store immediately prevents new transactions from being
// started. Specifically, it causes Begin() to return ErrDataStoreClosed.
//
// The behavior of any other read or write operation on a closed data-store
// is undefined.
//
// If there are any transactions in progress, Close() blocks until they are
// committed or rolled back.
func (ds *dataStore) Close() error {
	if err := ds.close(); err != nil {
		return err
	}

	<-ds.done

	return nil
}

// close marks the data-store as closed, and closes the done channel if there
// are no pending transactions.
func (ds *dataStore) close() error {
	ds.m.Lock()
	defer ds.m.Unlock()

	if ds.closed {
		return persistence.ErrDataStoreClosed
	}

	ds.closed = true
	ds.checkIfDone()

	return nil
}

// releaseTx decrements the transaction count and checks if the data-store is
// done.
func (ds *dataStore) releaseTx() {
	ds.m.Lock()
	defer ds.m.Unlock()

	ds.txns--
	ds.checkIfDone()
}

// checkIfDone closes the database and the done channel if the data-store is
// closed and there are no remaining transactions.
func (ds *dataStore) checkIfDone() {
	if ds.closed && ds.txns == 0 {
		ds.release(string(ds.appKey))
		close(ds.done)
	}
}
