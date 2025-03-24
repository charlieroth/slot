package userbus

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Storer interface {
	Create(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, user User) error
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, userID uuid.UUID) (User, error)
	Query(ctx context.Context, filter QueryFilter) ([]User, error)
}

type Business struct {
	logger *zerolog.Logger
	storer Storer
}

func NewBusiness(logger *zerolog.Logger, storer Storer) *Business {
	return &Business{
		logger: logger,
		storer: storer,
	}
}
