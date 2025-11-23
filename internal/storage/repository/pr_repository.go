package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"revass/internal/model"
	"revass/internal/storage"
)

const (
	PRStatusOpen   = "OPEN"
	PRStatusMerged = "MERGED"
)

type PRRepository interface {
	CreatePR(id string, name string, authorID string) error
	AssignReviewer(prID string, reviewerID string) error
	RemoveReviewer(prID string, reviewerID string) error
	GetByID(id string) (*model.PullRequest, error)
	HasPR(id string) error
	Merge(id string) error
}

type prRepository struct {
	db *sql.DB
}

func NewPRRepository(db *sql.DB) PRRepository {
	return &prRepository{db: db}
}

func (rep *prRepository) CreatePR(id string, name string, authorID string) error {
	const method = "CreatePR"

	_, err := rep.db.Exec("INSERT INTO pull_request (id, name, author_id) VALUES ($1, $2, $3) RETURNING id", id, name, authorID)

	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	return nil
}

func (rep *prRepository) AssignReviewer(prID string, reviewerID string) error {
	const method = "AssignReviewer"

	_, err := rep.db.Exec("INSERT INTO pr_reviewer (pr_id, reviewer_id) VALUES ($1, $2)", prID, reviewerID)

	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	return nil
}

func (rep *prRepository) RemoveReviewer(prID string, reviewerID string) error {
	const method = "RemoveReviewer"

	_, err := rep.db.Exec("DELETE FROM pr_reviewer WHERE pr_id = $1 AND reviewer_id = $2", prID, reviewerID)

	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	return nil
}

func (rep *prRepository) GetByID(id string) (*model.PullRequest, error) {
	const method = "GetByID"

	pr := new(model.PullRequest)

	err := rep.db.QueryRow(`
		SELECT 
		pull_request.id, pull_request.name, pull_request.author_id, pull_request.status, pull_request.merged_at
		FROM pull_request
		WHERE pull_request.id = $1;
	`, id).Scan(&pr.PullRequestID, &pr.PullRequestName, &pr.AuthorID, &pr.Status, &pr.MergedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", method, storage.ErrEntityNotFound)
		}

		return nil, fmt.Errorf("%s: %w", method, err)
	}

	rows, err := rep.db.Query(`
		SELECT 
		pr_reviewer.reviewer_id
		FROM pull_request
		JOIN pr_reviewer ON pull_request.id = pr_reviewer.pr_id
		WHERE pull_request.id = $1;
	`, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	for rows.Next() {
		var reviewerID string
		err = rows.Scan(&reviewerID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", method, err)
		}

		pr.AssignedReviewers = append(pr.AssignedReviewers, reviewerID)
	}

	return pr, nil
}

func (rep *prRepository) HasPR(id string) error {
	const method = "GetByID"

	var name string
	err := rep.db.QueryRow("SELECT name FROM pull_request WHERE id = $1", id).Scan(&name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", method, storage.ErrEntityNotFound)
		}

		return fmt.Errorf("%s: %w", method, err)
	}

	return nil
}

func (rep *prRepository) Merge(id string) error {
	const method = "Merge"

	_, err := rep.db.Exec("UPDATE pull_request SET status = 'MERGED', merged_at = CURRENT_TIMESTAMP WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	return nil
}
