package userbus

import "github.com/google/uuid"

type QueryFilter struct {
	ID       *uuid.UUID
	Email    *string
	Name     *string
	Phone    *string
	UserType *string
	TimeZone *string
}
