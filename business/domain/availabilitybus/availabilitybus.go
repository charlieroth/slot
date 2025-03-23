package availabilitybus

import (
	"context"

	"github.com/rs/zerolog"
)

type Storer interface {
	Query(ctx context.Context) error
	Update(ctx context.Context) error
}

type Business struct {
	logger *zerolog.Logger
	storer Storer
}

func NewBusiness(storer Storer, logger *zerolog.Logger) *Business {
	return &Business{
		storer: storer,
		logger: logger,
	}
}
