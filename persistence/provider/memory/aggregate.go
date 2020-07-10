package memory

import (
	"context"

	"github.com/dogmatiq/infix/persistence"
)

// LoadAggregateMetaData loads the meta-data for an aggregate instance.
//
// hk is the aggregate handler's identity key, id is the instance ID.
func (ds *dataStore) LoadAggregateMetaData(
	ctx context.Context,
	hk, id string,
) (persistence.AggregateMetaData, error) {
	ds.db.mutex.RLock()
	defer ds.db.mutex.RUnlock()

	key := instanceKey{hk, id}
	if md, ok := ds.db.aggregate.metadata[key]; ok {
		return md, nil
	}

	return persistence.AggregateMetaData{
		HandlerKey: hk,
		InstanceID: id,
	}, nil
}

// aggregateDatabase contains aggregate related data.
type aggregateDatabase struct {
	metadata map[instanceKey]persistence.AggregateMetaData
}

// VisitSaveAggregateMetaData returns an error if a "SaveAggregateMetaData"
// operation can not be applied to the database.
func (v *validator) VisitSaveAggregateMetaData(
	_ context.Context,
	op persistence.SaveAggregateMetaData,
) error {
	new := op.MetaData
	key := instanceKey{new.HandlerKey, new.InstanceID}
	old := v.db.aggregate.metadata[key]

	if new.Revision == old.Revision {
		return nil
	}

	return persistence.ConflictError{
		Cause: op,
	}
}

// VisitSaveAggregateMetaData applies the changes in a "SaveAggregateMetaData"
// operation to the database.
func (c *committer) VisitSaveAggregateMetaData(
	_ context.Context,
	op persistence.SaveAggregateMetaData,
) error {
	md := op.MetaData
	key := instanceKey{md.HandlerKey, md.InstanceID}

	if c.db.aggregate.metadata == nil {
		c.db.aggregate.metadata = map[instanceKey]persistence.AggregateMetaData{}
	}

	md.Revision++
	c.db.aggregate.metadata[key] = md

	return nil
}
