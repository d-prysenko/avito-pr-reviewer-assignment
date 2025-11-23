package response

import (
	"net/http"
	"revass/internal/handler"
	"revass/internal/model"
)

const (
	ErrorCodePRExists = "PR_EXISTS"
	ErrorCodePRMerged = "PR_MERGED"
)

type PRCreateResponse struct {
	PR *model.PullRequest `json:"pr"`
}

type PRReassignResponse struct {
	PR         *model.PullRequest `json:"pr"`
	ReplacedBy string             `json:"replaced_by"`
}

func MakePRAlreadyExistsReponse(w http.ResponseWriter, prID string) {
	handler.MakeErrorResponse(
		w,
		handler.ErrorResponse{
			Code:    ErrorCodePRExists,
			Message: "PR '" + prID + "' already exists",
		},
		http.StatusConflict,
	)
}

func MakePRMergedReponse(w http.ResponseWriter) {
	handler.MakeErrorResponse(
		w,
		handler.ErrorResponse{
			Code:    ErrorCodePRMerged,
			Message: "cannot reassign on merged PR",
		},
		http.StatusConflict,
	)
}
