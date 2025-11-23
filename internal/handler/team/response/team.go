package response

import (
	"net/http"
	"revass/internal/handler"
	"revass/internal/model"
)

const (
	ErrorCodeTeamExists = "TEAM_EXISTS"
	ErrorCodeUserExists = "USER_EXISTS"
)

type TeamAddResponse struct {
	Team model.Team `json:"team"`
}

func MakeUserAlreadyExistsReponse(w http.ResponseWriter, userID string) {
	handler.MakeErrorResponse(
		w,
		handler.ErrorResponse{
			Code:    ErrorCodeUserExists,
			Message: "User '" + userID + "' already exists",
		},
		http.StatusBadRequest,
	)
}

func MakeTeamAlreadyExistsReponse(w http.ResponseWriter, teamID string) {
	handler.MakeErrorResponse(
		w,
		handler.ErrorResponse{
			Code:    ErrorCodeTeamExists,
			Message: "Team '" + teamID + "' already exists",
		},
		http.StatusBadRequest,
	)
}
