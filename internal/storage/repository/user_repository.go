package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"revass/internal/model"
	"revass/internal/storage"
)

type UserRepository interface {
	CreateUser(userID string, username string, isActive bool) error
	GetUserByID(userID string) (*model.User, error)
	SetIsActive(userID string, isActive bool) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (rep *userRepository) CreateUser(userID string, username string, isActive bool) error {
	const method = "CreateUser"

	_, err := rep.db.Exec("INSERT INTO users (id, username, is_active) VALUES ($1, $2, $3)", userID, username, isActive)

	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	return nil
}

func (rep *userRepository) GetUserByID(userID string) (*model.User, error) {
	const method = "GetUserByID"

	row := rep.db.QueryRow("SELECT * FROM users WHERE id = $1", userID)

	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", method, storage.ErrEntityNotFound)
		}

		return nil, fmt.Errorf("%s: %w", method, err)
	}

	return user, nil
}

func (rep *userRepository) SetIsActive(userID string, isActive bool) error {
	const method = "SetIsActive"

	_, err := rep.db.Exec("UPDATE users SET is_active = $1 WHERE id = $2", isActive, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	return nil
}

func scanUser(row *sql.Row) (*model.User, error) {
	user := new(model.User)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.IsActive,
		&user.Team,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
