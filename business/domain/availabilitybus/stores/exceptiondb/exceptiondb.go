package exceptiondb

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Store struct {
	logger *zerolog.Logger
	db     *pgxpool.Pool
}

func NewStore(logger *zerolog.Logger, db *pgxpool.Pool) *Store {
	return &Store{
		logger: logger,
		db:     db,
	}
}

func (s *Store) Query(ctx context.Context) error {
	return nil
}

func (s *Store) Update(ctx context.Context) error {
	return nil
}
