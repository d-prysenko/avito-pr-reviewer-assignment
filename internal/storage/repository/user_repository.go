package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"revass/internal/model"
	"revass/internal/storage"
)

type UserRepository interface {
	CreateUser(userStrID string, username string, isActive bool) (int64, error)
	GetUserByID(userID int) (*model.User, error)
	SetIsActive(userID int, isActive bool) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (rep *userRepository) CreateUser(userStrID string, username string, isActive bool) (int64, error) {
	const method = "CreateUser"

	var userID int64

	err := rep.db.QueryRow("INSERT INTO users (user_str_id, username, is_active) VALUES ($1, $2, $3) RETURNING id", userStrID, username, isActive).Scan(&userID)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", method, err)
	}

	return userID, nil
}

func (rep *userRepository) GetUserByID(userID int) (*model.User, error) {
	const method = "GetUserByID"

	row := rep.db.QueryRow("SELECT * FROM users WHERE user_id = $1", userID)

	user, err := scanUser(row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", method, storage.ErrEntityNotFound)
		}

		return nil, fmt.Errorf("%s: %w", method, err)
	}

	return user, nil
}

func (rep *userRepository) SetIsActive(userID int, isActive bool) error {
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
		&user.StrID,
		&user.Username,
		&user.IsActive,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
