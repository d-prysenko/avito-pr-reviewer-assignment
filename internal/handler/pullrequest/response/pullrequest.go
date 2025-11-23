package response

import (
	"net/http"
	"revass/internal/handler"
	"revass/internal/model"
)

const (
	ERR_CODE_PR_EXISTS = "PR_EXISTS"
	ERR_CODE_PR_MERGED = "PR_MERGED"
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
			Code:    ERR_CODE_PR_EXISTS,
			Message: "PR '" + prID + "' already exists",
		},
		http.StatusConflict,
	)
}

func MakePRMergedReponse(w http.ResponseWriter) {
	handler.MakeErrorResponse(
		w,
		handler.ErrorResponse{
			Code:    ERR_CODE_PR_MERGED,
			Message: "cannot reassign on merged PR",
		},
		http.StatusConflict,
	)
}
