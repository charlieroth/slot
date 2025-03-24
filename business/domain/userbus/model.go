package userbus

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Name      string
	Phone     *string
	UserType  string
	TimeZone  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewUser struct {
	Email    string
	Name     string
	Phone    *string
	UserType string
	TimeZone string
}

type UpdateUser struct {
	Email    *string
	Name     *string
	Phone    *string
	UserType *string
	TimeZone *string
}
