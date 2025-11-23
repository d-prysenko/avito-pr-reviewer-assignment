package service

import (
	"errors"
	"fmt"
	"math/rand"
	"revass/internal/model"
	"revass/internal/storage"
	"revass/internal/storage/repository"
	"slices"
	"time"
)

var (
	ErrPRMerged = errors.New("PR merged")
)

type PRManager interface {
	Create(id string, name string, authorID string) (*model.PullRequest, error)
	Merge(id string) (*model.PullRequest, error)
	Reassign(prID string, oldReviewerID string) (*model.PullRequest, string, error)
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

func (prm *prManager) Reassign(prID string, oldReviewerID string) (*model.PullRequest, string, error) {
	const method = "Reassign"

	pr, err := prm.prRep.GetByID(prID)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", method, err)
	}

	if pr.Status == repository.PRStatusMerged {
		return nil, "", fmt.Errorf("%s: %w", method, ErrPRMerged)
	}

	oldReviewer, err := prm.userRep.GetUserByID(oldReviewerID)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", method, err)
	}

	members, err := prm.teamRep.GetActiveTeamMembersExcludingUser(oldReviewer.Team, pr.AuthorID)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", method, err)
	}

	otherReviewer := ""

	for _, rID := range pr.AssignedReviewers {
		if rID != oldReviewerID {
			otherReviewer = rID
			break
		}
	}

	members = slices.DeleteFunc(members, func(member *model.TeamMember) bool {
		return member.UserID == oldReviewerID || member.UserID == otherReviewer
	})

	err = prm.prRep.RemoveReviewer(prID, oldReviewerID)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", method, err)
	}

	if len(members) == 0 {
		return nil, "", nil
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	i := r.Intn(len(members))

	err = prm.prRep.AssignReviewer(prID, members[i].UserID)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", method, err)
	}

	pr, err = prm.prRep.GetByID(prID)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", method, err)
	}

	return pr, members[i].UserID, nil
}
