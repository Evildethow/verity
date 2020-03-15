package postgres_test

import (
	"context"
	"database/sql"
	"time"

	"github.com/dogmatiq/infix/envelope"
	"github.com/dogmatiq/infix/internal/testing/sqltest"
	"github.com/dogmatiq/infix/internal/testing/streamtest"
	"github.com/dogmatiq/infix/internal/x/sqlx"
	infixsql "github.com/dogmatiq/infix/persistence/provider/sql"
	. "github.com/dogmatiq/infix/persistence/provider/sql/driver/postgres"
	"github.com/dogmatiq/linger/backoff"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type StreamDriver (standard test suite)", func() {
	var db *sql.DB

	streamtest.Declare(
		func(ctx context.Context, in streamtest.In) streamtest.Out {
			db = sqltest.Open("postgres")

			err := DropSchema(ctx, db)
			Expect(err).ShouldNot(HaveOccurred())

			err = CreateSchema(ctx, db)
			Expect(err).ShouldNot(HaveOccurred())

			stream := &infixsql.Stream{
				ApplicationKey:  "<app-key>",
				DB:              db,
				Driver:          StreamDriver{},
				Marshaler:       in.Marshaler,
				BackoffStrategy: backoff.Constant(10 * time.Millisecond),
			}

			return streamtest.Out{
				Stream: stream,
				Append: func(ctx context.Context, envelopes ...*envelope.Envelope) {
					tx := sqlx.Begin(ctx, db)
					defer tx.Rollback()

					_, err := stream.Append(
						ctx,
						tx,
						envelopes...,
					)
					Expect(err).ShouldNot(HaveOccurred())

					sqlx.Commit(tx)
				},
			}
		},
		func() {
			if db != nil {
				db.Close()
			}
		},
	)
})
