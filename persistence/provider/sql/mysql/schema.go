package mysql

import (
	"context"
	"database/sql"

	"github.com/dogmatiq/infix/internal/x/sqlx"
)

// CreateSchema creates the schema elements required by the MySQL driver.
func CreateSchema(ctx context.Context, db *sql.DB) (err error) {
	defer sqlx.Recover(&err)

	createEventStoreSchema(ctx, db)

	return nil
}

// DropSchema drops the schema elements required by the MySQL driver.
func DropSchema(ctx context.Context, db *sql.DB) (err error) {
	defer sqlx.Recover(&err)

	sqlx.Exec(ctx, db, `DROP TABLE IF EXISTS event_offset`)
	sqlx.Exec(ctx, db, `DROP TABLE IF EXISTS event`)
	sqlx.Exec(ctx, db, `DROP TABLE IF EXISTS event_filter`)
	sqlx.Exec(ctx, db, `DROP TABLE IF EXISTS event_filter_name`)

	return nil
}

// createEventStoreSchema creates the schema elements required by the event
// store subsystem.
func createEventStoreSchema(ctx context.Context, db *sql.DB) {
	sqlx.Exec(
		ctx,
		db,
		`CREATE TABLE event_offset (
			source_app_key VARBINARY(255) NOT NULL PRIMARY KEY,
			next_offset    BIGINT NOT NULL
		) ENGINE=InnoDB`,
	)

	sqlx.Exec(
		ctx,
		db,
		`CREATE TABLE event (
			offset              BIGINT UNSIGNED NOT NULL,
			message_id          VARBINARY(255) NOT NULL,
			causation_id        VARBINARY(255) NOT NULL,
			correlation_id      VARBINARY(255) NOT NULL,
			source_app_name     VARBINARY(255) NOT NULL,
			source_app_key      VARBINARY(255) NOT NULL,
			source_handler_name VARBINARY(255) NOT NULL,
			source_handler_key  VARBINARY(255) NOT NULL,
			source_instance_id  VARBINARY(255) NOT NULL,
			created_at          VARBINARY(255) NOT NULL,
			portable_name       VARBINARY(255) NOT NULL,
			media_type          VARBINARY(255) NOT NULL,
			data                LONGBLOB NOT NULL,

			PRIMARY KEY (source_app_key, offset),
			INDEX eventstore_query (
				source_app_key,
				portable_name,
				offset,
				source_handler_key,
				source_instance_id
			)
		) ENGINE=InnoDB ROW_FORMAT=COMPRESSED KEY_BLOCK_SIZE=4`,
	)

	sqlx.Exec(
		ctx,
		db,
		`CREATE TABLE event_filter (
			id      SERIAL PRIMARY KEY,
			app_key VARBINARY(255) NOT NULL UNIQUE
		) ENGINE=InnoDB`,
	)

	sqlx.Exec(
		ctx,
		db,
		`CREATE TABLE event_filter_name (
			filter_id     BIGINT UNSIGNED NOT NULL REFERENCES event_filter (id) ON DELETE CASCADE,
			portable_name VARBINARY(255) NOT NULL,

			PRIMARY KEY (filter_id, portable_name)
		) ENGINE=InnoDB`,
	)
}