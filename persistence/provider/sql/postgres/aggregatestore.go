package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dogmatiq/infix/persistence/subsystem/aggregatestore"
)

// InsertAggregateRevision inserts an aggregate revision for an aggregate
// instance.
//
// It returns false if the row already exists.
func (driver) InsertAggregateRevision(
	ctx context.Context,
	tx *sql.Tx,
	ak, hk, id string,
) (bool, error) {
	return false, errors.New("not implemented")
}

// UpdateAggregateRevision increments an aggregate isntance's revision by 1.
//
// It returns false if the row does not exist or rev is not current.
func (driver) UpdateAggregateRevision(
	ctx context.Context,
	tx *sql.Tx,
	ak, hk, id string,
	rev aggregatestore.Revision,
) (bool, error) {
	return false, errors.New("not implemented")
}

// SelectAggregateRevision selects an aggregate instance's revision.
func (driver) SelectAggregateRevision(
	ctx context.Context,
	db *sql.DB,
	ak, hk, id string,
) (aggregatestore.Revision, error) {
	row := db.QueryRowContext(
		ctx,
		`SELECT
			r.revision
		FROM infix.aggregate_revision AS r
		WHERE app_key = $1
		AND handler_key = $2
		AND instance_id = $3`,
		ak,
		hk,
		id,
	)

	var rev aggregatestore.Revision
	err := row.Scan(&rev)

	if err == sql.ErrNoRows {
		return 0, nil
	}

	return rev, err
}
