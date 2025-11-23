package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"revass/internal/model"
	"revass/internal/storage"

	"github.com/lib/pq"
)

type TeamRepository interface {
	CreateTeam(teamName string) (int64, error)
	GetTeamIDByName(teamName string) (int64, error)
	AddTeamAndUsers(team model.Team) error
	GetTeamMembersByID(teamID int64) ([]*model.TeamMember, error)
	GetActiveTeamMembersExcludingUser(teamID int64, userID string) ([]*model.TeamMember, error)
}

type teamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (rep *teamRepository) CreateTeam(teamName string) (int64, error) {
	const method = "CreateTeam"

	var teamID int64

	err := rep.db.QueryRow("INSERT INTO team (name) VALUES ($1) RETURNING id", teamName).Scan(&teamID)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", method, err)
	}

	return teamID, nil
}

func (rep *teamRepository) GetTeamIDByName(teamName string) (int64, error) {
	const method = "GetTeamIDByName"

	var teamID int64

	err := rep.db.QueryRow("SELECT id FROM team WHERE name = $1", teamName).Scan(&teamID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("%s: %w", method, storage.ErrEntityNotFound)
		}

		return 0, fmt.Errorf("%s: %w", method, err)
	}

	return teamID, nil
}

func (rep *teamRepository) AddTeamAndUsers(team model.Team) error {
	const method = "AddTeamAndUsers"

	_, err := rep.GetTeamIDByName(team.Name)
	if err == nil {
		return fmt.Errorf("%s: %w", method, &storage.ErrEntityExists{ID: team.Name, Err: storage.ErrTeamExists})
	}
	if !errors.Is(err, storage.ErrEntityNotFound) {
		return fmt.Errorf("%s: %w", method, err)
	}

	tx, err := rep.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	var teamID int64
	err = tx.QueryRow("INSERT INTO team (name) VALUES ($1) RETURNING id", team.Name).Scan(&teamID)
	if err != nil {
		tx.Rollback()

		return fmt.Errorf("%s: %w", method, err)
	}

	for _, member := range team.Members {
		_, err = tx.Exec("INSERT INTO users (id, username, is_active, team_id) VALUES ($1, $2, $3, $4)", member.UserID, member.Username, member.IsActive, teamID)
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation" {
				tx.Rollback()

				return fmt.Errorf("%s: %w", method, &storage.ErrEntityExists{ID: member.UserID, Err: storage.ErrUserExists})
			}

			tx.Rollback()
			
			return fmt.Errorf("%s: %w", method, err)
		}
	}

	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("%s: %w", method, err)
	}

	return nil
}

func (rep *teamRepository) GetTeamMembersByID(teamID int64) ([]*model.TeamMember, error) {
	const method = "GetTeamMembersByID"

	var members []*model.TeamMember

	rows, err := rep.db.Query(`
		SELECT 
		users.id, users.username, users.is_active
		FROM users
		WHERE users.team_id = $1;
	`, teamID)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	for rows.Next() {
		member, err := scanTeamMember(rows)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", method, err)
		}

		members = append(members, member)
	}

	return members, nil
}

func (rep *teamRepository) GetActiveTeamMembersExcludingUser(teamID int64, userID string) ([]*model.TeamMember, error) {
	const method = "GetActiveTeamMembersExcludingUser"

	var members []*model.TeamMember

	rows, err := rep.db.Query(`
		SELECT 
		users.id, users.username, users.is_active
		FROM users
		WHERE users.team_id = $1 AND users.id != $2 AND users.is_active = TRUE;
	`, teamID, userID)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", method, err)
	}

	for rows.Next() {
		member, err := scanTeamMember(rows)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", method, err)
		}

		members = append(members, member)
	}

	return members, nil
}

func scanTeamMember(rows *sql.Rows) (*model.TeamMember, error) {
	teamMember := new(model.TeamMember)

	err := rows.Scan(
		&teamMember.UserID,
		&teamMember.Username,
		&teamMember.IsActive,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrEntityNotFound
		}

		return nil, err
	}

	return teamMember, nil
}
