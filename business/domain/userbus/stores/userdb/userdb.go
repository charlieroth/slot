package userdb

import (
	"bytes"
	"context"
	"fmt"

	"github.com/charlieroth/slot/business/domain/userbus"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Store struct {
	logger *zerolog.Logger
	db     *pgxpool.Pool
}

func NewStore(logger *zerolog.Logger, db *pgxpool.Pool) *Store {
	return &Store{
		logger: logger,
		db:     db,
	}
}

func (s *Store) Create(ctx context.Context, user userbus.User) error {
	const q = `
	INSERT INTO users
		(id, email, name, phone, user_type, time_zone, created_at, updated_at)
	VALUES
		(@id, @email, @name, @phone, @user_type, @time_zone, @created_at, @updated_at)
	`
	_, err := s.db.Exec(ctx, q, pgx.NamedArgs{
		"id":         user.ID,
		"email":      user.Email,
		"name":       user.Name,
		"phone":      user.Phone,
		"user_type":  user.UserType,
		"time_zone":  user.TimeZone,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (s *Store) Update(ctx context.Context, user userbus.User) error {
	const q = `
	UPDATE
		users
	SET
		email = @email,
		name = @name,
		phone = @phone,
		user_type = @user_type,
		time_zone = @time_zone,
		updated_at = @updated_at
	WHERE id = @id
	`
	_, err := s.db.Exec(ctx, q, pgx.NamedArgs{
		"id":         user.ID,
		"email":      user.Email,
		"name":       user.Name,
		"phone":      user.Phone,
		"user_type":  user.UserType,
		"time_zone":  user.TimeZone,
		"updated_at": user.UpdatedAt,
	})
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return nil
}

func (s *Store) Delete(ctx context.Context, user userbus.User) error {
	const q = `
	DELETE FROM users
	WHERE id = @id
	`
	_, err := s.db.Exec(ctx, q, pgx.NamedArgs{
		"id": user.ID,
	})
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}

func (s *Store) Count(ctx context.Context, filter userbus.QueryFilter) (int, error) {
	const q = `
	SELECT count(1)
	FROM users
	`

	rows, err := s.db.Query(ctx, q, pgx.NamedArgs{
		"id":        filter.ID,
		"email":     filter.Email,
		"name":      filter.Name,
		"phone":     filter.Phone,
		"user_type": filter.UserType,
		"time_zone": filter.TimeZone,
	})
	if err != nil {
		return 0, fmt.Errorf("count users: %w", err)
	}

	count, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[struct {
		Count int `db:"count"`
	}])
	if err != nil {
		return 0, fmt.Errorf("collect one row: %w", err)
	}

	return count.Count, nil
}

func (s *Store) QueryByID(ctx context.Context, userID uuid.UUID) (userbus.User, error) {
	const q = `
	SELECT id, email, name, phone, user_type, time_zone, created_at, updated_at
	FROM users
	WHERE id = @id
	`

	rows, err := s.db.Query(ctx, q, pgx.NamedArgs{
		"id": userID,
	})
	if err != nil {
		return userbus.User{}, fmt.Errorf("query user by id: %w", err)
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return userbus.User{}, fmt.Errorf("collect one row: %w", err)
	}

	return toBusUser(user), nil
}

func (s *Store) Query(ctx context.Context, filter userbus.QueryFilter) ([]userbus.User, error) {
	data := make(map[string]any)

	const q = `
	SELECT id, email, name, phone, user_type, time_zone, created_at, updated_at
	FROM users
	`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	rows, err := s.db.Query(ctx, buf.String(), data)
	if err != nil {
		return nil, fmt.Errorf("query users: %w", err)
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, fmt.Errorf("collect rows: %w", err)
	}

	return toBusUsers(users), nil
}
