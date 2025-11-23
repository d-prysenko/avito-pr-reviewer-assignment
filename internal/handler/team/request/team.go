package request

import "revass/internal/model"

type TeamAddRequest struct {
	Name    string       `json:"team_name" validate:"required,max=128"`
	Members []TeamMember `json:"members" validate:"required"`
}

type TeamMember struct {
	UserID   string `json:"user_id" validate:"required,max=128"`
	Username string `json:"username" validate:"required,max=128"`
	IsActive bool   `json:"is_active" validate:"required,boolean"`
}

func (t *TeamAddRequest) ToTeamModel() model.Team {
	var team model.Team

	team.Name = t.Name
	team.Members = make([]model.TeamMember, len(t.Members))

	for i := range t.Members {
		team.Members[i] = model.TeamMember{
			UserID:   t.Members[i].UserID,
			Username: t.Members[i].Username,
			IsActive: t.Members[i].IsActive,
		}
	}

	return team
}
