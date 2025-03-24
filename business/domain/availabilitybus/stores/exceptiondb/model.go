package exceptiondb

import (
	"time"

	"github.com/google/uuid"
)

type AvailabilityException struct {
	ID            uuid.UUID  `db:"id"`
	MakerID       uuid.UUID  `db:"maker_id"`
	ExceptionDate time.Time  `db:"exception_date"` // only date component is used (time is ignored)
	StartTime     *time.Time `db:"start_time"`
	EndTime       *time.Time `db:"end_time"`
	IsAvailable   bool       `db:"is_available"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
}

// Example for setting the value of an exception date:
// exceptionDate = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
