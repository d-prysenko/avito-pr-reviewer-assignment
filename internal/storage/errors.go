package storage

import (
	"errors"
	"fmt"
)

var (
	ErrEntityNotFound = errors.New("entity not found")
	ErrUserExists     = errors.New("user already exists")
	ErrTeamExists     = errors.New("team already exists")
	ErrPRExists       = errors.New("pull request already exists")
)

type ErrEntityExists struct {
	ID  string
	Err error
}

func (e *ErrEntityExists) Error() string {
	return fmt.Sprintf("entity with id '%s' already exists", e.ID)
}

func (e *ErrEntityExists) Unwrap() error {
	return e.Err
}
