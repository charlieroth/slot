package userdb

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	Phone     *string   `db:"phone"`
	UserType  string    `db:"user_type"`
	TimeZone  string    `db:"time_zone"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const (
	UserTypeMaker    = "maker"
	UserTypeCustomer = "customer"
	UserTypeAdmin    = "admin"
)
