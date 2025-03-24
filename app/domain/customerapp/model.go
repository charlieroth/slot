package customerapp

import (
	"fmt"
	"time"

	"github.com/charlieroth/slot/app/sdk/errs"
	"github.com/charlieroth/slot/business/domain/userbus"
)

type User struct {
	ID        string  `json:"id"`
	Email     string  `json:"email"`
	Name      string  `json:"name"`
	Phone     *string `json:"phone"`
	UserType  string  `json:"user_type"`
	TimeZone  string  `json:"time_zone"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func toBusUser(bus userbus.User) User {
	return User{
		ID:        bus.ID.String(),
		Email:     bus.Email,
		Name:      bus.Name,
		Phone:     bus.Phone,
		UserType:  bus.UserType,
		TimeZone:  bus.TimeZone,
		CreatedAt: bus.CreatedAt.Format(time.RFC3339),
		UpdatedAt: bus.UpdatedAt.Format(time.RFC3339),
	}
}

func toBusUsers(bus []userbus.User) []User {
	users := make([]User, len(bus))
	for i, user := range bus {
		users[i] = toBusUser(user)
	}
	return users
}

type NewUser struct {
	Email    string  `json:"email" validate:"required,email"`
	Name     string  `json:"name" validate:"required"`
	Phone    *string `json:"phone"`
	UserType string  `json:"user_type" validate:"required"`
	TimeZone string  `json:"time_zone" validate:"required"`
}

func (app NewUser) Validate() error {
	if err := errs.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}

func toBusNewUser(app NewUser) (userbus.NewUser, error) {
	// TODO: handle parsing to business layer types once implemented

	bus := userbus.NewUser{
		Email:    app.Email,
		Name:     app.Name,
		Phone:    app.Phone,
		UserType: app.UserType,
		TimeZone: app.TimeZone,
	}

	return bus, nil
}

type UpdateUser struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	Name     *string `json:"name"`
	Phone    *string `json:"phone"`
	UserType *string `json:"user_type"`
	TimeZone *string `json:"time_zone"`
}

func (app UpdateUser) Validate() error {
	if err := errs.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}

func toBusUpdateUser(app UpdateUser) (userbus.UpdateUser, error) {
	// TODO: handle parsing to business layer types once implemented

	bus := userbus.UpdateUser{
		Email:    app.Email,
		Name:     app.Name,
		Phone:    app.Phone,
		UserType: app.UserType,
		TimeZone: app.TimeZone,
	}

	return bus, nil
}
