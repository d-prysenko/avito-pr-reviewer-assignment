package service

import (
	"fmt"
	"revass/internal/model"
	"revass/internal/storage/repository"
)

type UserManager interface {
	SetIsActive(userID string, isActive bool) (*model.User, error)
	GetReview(userID string) ([]*model.UserPullRequest, error)
}

type userManager struct {
	userRep repository.UserRepository
}

func NewUserManager(userRep repository.UserRepository) UserManager {
	return &userManager{userRep: userRep}
}

func (um *userManager) SetIsActive(userID string, isActive bool) (*model.User, error) {
	const method = "SetIsActive"

	_, err := um.userRep.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	err = um.userRep.SetIsActive(userID, isActive)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	user, err := um.userRep.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	return user, nil
}

func (um *userManager) GetReview(userID string) ([]*model.UserPullRequest, error) {
	const method = "GetReview"

	_, err := um.userRep.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	prs, err := um.userRep.GetReview(userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	return prs, nil
}
