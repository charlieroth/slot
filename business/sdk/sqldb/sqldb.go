package sqldb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/charlieroth/slot/foundation/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// lib/pq errorCodeNames
// https://github.com/lib/pq/blob/master/error.go
const (
	uniqueViolation = "23505"
	undefinedTable  = "42P01"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound          = sql.ErrNoRows
	ErrDBDuplicatedEntry = errors.New("duplicated entry")
	ErrDBUndefinedTable  = errors.New("undefined table")
)

// Open knows how to open a database connection based on the configuration.
func Open(cfg config.DBConfig) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	pool.Config().MinConns = cfg.MinConns
	pool.Config().MaxConns = cfg.MaxConns

	return pool, nil
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func StatusCheck(ctx context.Context, pool *pgxpool.Pool) error {
	// If the user doesn't give us a deadline, set 1 second.
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second)
		defer cancel()
	}

	for attempts := 1; ; attempts++ {
		if err := pool.Ping(ctx); err == nil {
			break
		}

		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Run simple query to check connectivity
	const q = `SELECT TRUE`
	var tmp bool
	return pool.QueryRow(ctx, q).Scan(&tmp)
}
