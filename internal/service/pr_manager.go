package service

import (
	"fmt"
	"math/rand"
	"revass/internal/model"
	"revass/internal/storage"
	"revass/internal/storage/repository"
	"time"
)

type PRManager interface {
	Create(id string, name string, authorID string) (*model.PullRequest, error)
	Merge(id string) (*model.PullRequest, error)
}

type prManager struct {
	prRep   repository.PRRepository
	userRep repository.UserRepository
	teamRep repository.TeamRepository
}

func NewPRManager(
	prRep repository.PRRepository,
	userRep repository.UserRepository,
	teamRep repository.TeamRepository,
) PRManager {
	return &prManager{
		prRep:   prRep,
		userRep: userRep,
		teamRep: teamRep,
	}
}

func (prm *prManager) Create(id string, name string, authorID string) (*model.PullRequest, error) {
	method := "Create"

	author, err := prm.userRep.GetUserByID(authorID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	err = prm.prRep.HasPR(id)
	if err == nil {
		return nil, fmt.Errorf("%s: %w", method, &storage.ErrEntityExists{ID: id, Err: storage.ErrPRExists})
	}

	err = prm.prRep.CreatePR(id, name, authorID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	members, err := prm.teamRep.GetActiveTeamMembersExcludingUser(author.Team, authorID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	membersCount := len(members)

	if membersCount >= 2 {
		i1, i2 := getDifferentRandomInts(membersCount)

		err = prm.prRep.AssignReviewer(id, members[i1].UserID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", method, err)
		}

		err = prm.prRep.AssignReviewer(id, members[i2].UserID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", method, err)
		}
	} else if membersCount == 1 {
		err = prm.prRep.AssignReviewer(id, members[0].UserID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", method, err)
		}
	}

	pr, err := prm.prRep.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	return pr, nil
}

func getDifferentRandomInts(max int) (i1 int, i2 int) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if max >= 2 {
		i1 = r.Intn(max)
		i2 = r.Intn(max - 1)

		if i2 >= i1 {
			i2++
		}
	}

	return i1, i2
}

func (prm *prManager) Merge(id string) (*model.PullRequest, error) {
	const method = "Merge"

	pr, err := prm.prRep.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	if pr.Status == repository.PRStatusMerged {
		return pr, nil
	}

	err = prm.prRep.Merge(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	pr, err = prm.prRep.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	return pr, nil
}
