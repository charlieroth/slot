package servicedb

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID              uuid.UUID `db:"id"`
	MakerID         uuid.UUID `db:"maker_id"`
	Name            string    `db:"name"`
	Description     *string   `db:"description"`
	DurationMinutes int       `db:"duration_minutes"`
	BufferMinutes   int       `db:"buffer_minutes"`
	Price           *float64  `db:"price"`
	IsActive        bool      `db:"is_active"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
