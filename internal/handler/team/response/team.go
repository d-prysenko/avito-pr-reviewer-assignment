package response

import (
	"net/http"
	"revass/internal/handler"
	"revass/internal/model"
)

const (
	ERR_CODE_TEAM_EXISTS = "TEAM_EXISTS"
	ERR_CODE_USER_EXISTS = "USER_EXISTS"
)

type TeamAddResponse struct {
	Team model.Team `json:"team"`
}

func MakeUserAlreadyExistsReponse(w http.ResponseWriter, userID string) {
	handler.MakeErrorResponse(
		w,
		handler.ErrorResponse{
			Code:    ERR_CODE_USER_EXISTS,
			Message: "User '" + userID + "' already exists",
		},
		http.StatusBadRequest,
	)
}

func MakeTeamAlreadyExistsReponse(w http.ResponseWriter, teamID string) {
	handler.MakeErrorResponse(
		w,
		handler.ErrorResponse{
			Code:    ERR_CODE_TEAM_EXISTS,
			Message: "Team '" + teamID + "' already exists",
		},
		http.StatusBadRequest,
	)
}
