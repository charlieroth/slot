package scheduledb

import (
	"time"

	"github.com/google/uuid"
)

type MakerSchedule struct {
	ID        uuid.UUID `db:"id"`
	MakerID   uuid.UUID `db:"maker_id"`
	DayOfWeek int       `db:"day_of_week"` // 0-6 for Monday-Sunday
	StartTime time.Time `db:"start_time"`  // Only the time component is used
	EndTime   time.Time `db:"end_time"`    // Only the time component is used
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Example for setting the value of a time:
// startTime := time.Date(1, 1, 1, hour, minute, second, 0, time.UTC)
// endTime := time.Date(1, 1, 1, hour, minute, second, 0, time.UTC)

// Example for setting the value of a time:
// startTime := time.Date(0001, time.January, 1, 9, 0, 0, 0, time.UTC) // 9:00 AM
// endTime := time.Date(0001, time.January, 1, 17, 0, 0, 0, time.UTC)  // 5:00 PM
