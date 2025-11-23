package service

import (
	"fmt"
	"revass/internal/model"
	"revass/internal/storage/repository"
)

type TeamManager interface {
	AddTeamWithMembers(team model.Team) error
}

type teamManager struct {
	userRep repository.UserRepository
	teamRep repository.TeamRepository
}

func NewTeamManager(userRep repository.UserRepository, teamRep repository.TeamRepository) TeamManager {
	return &teamManager{userRep: userRep, teamRep: teamRep}
}

func (tm *teamManager) AddTeamWithMembers(team model.Team) error {
	method := "AddTeamWithMembers"

	err := tm.teamRep.AddTeamAndUsers(team)
	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	return nil
}
