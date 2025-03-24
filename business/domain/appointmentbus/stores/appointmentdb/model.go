package appointmentdb

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID         uuid.UUID `db:"id"`
	MakerID    uuid.UUID `db:"maker_id"`
	CustomerID uuid.UUID `db:"customer_id"`
	ServiceID  uuid.UUID `db:"service_id"`
	StartTime  time.Time `db:"start_time"`
	EndTime    time.Time `db:"end_time"`
	Status     string    `db:"status"` // valid values: pending, confirmed, cancelled, completed, no_show
	Notes      string    `db:"notes"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

const (
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusCancelled = "cancelled"
	StatusCompleted = "completed"
	StatusNoShow    = "no_show"
)
