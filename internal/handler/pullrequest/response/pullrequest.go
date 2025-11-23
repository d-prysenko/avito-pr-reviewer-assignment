package response

import (
	"net/http"
	"revass/internal/handler"
	"revass/internal/model"
)

const (
	ERR_CODE_PR_EXISTS = "PR_EXISTS"
)


type PRCreateResponse struct {
	PR *model.PullRequest `json:"pr"`
}

func MakePRAlreadyExistsReponse(w http.ResponseWriter, prID string) {
	handler.MakeErrorResponse(
		w,
		handler.ErrorResponse{
			Code:    ERR_CODE_PR_EXISTS,
			Message: "PR '" + prID + "' already exists",
		},
		http.StatusConflict,
	)
}