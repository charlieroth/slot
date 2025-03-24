package userdb

import (
	"time"

	"github.com/charlieroth/slot/business/domain/userbus"
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

func toBusUser(user User) userbus.User {
	return userbus.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Phone:     user.Phone,
		UserType:  user.UserType,
		TimeZone:  user.TimeZone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toBusUsers(users []User) []userbus.User {
	busUsers := make([]userbus.User, len(users))
	for i, user := range users {
		busUsers[i] = toBusUser(user)
	}
	return busUsers
}
