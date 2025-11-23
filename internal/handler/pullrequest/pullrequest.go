package pullrequest

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"revass/internal/handler"
	"revass/internal/handler/pullrequest/request"
	"revass/internal/handler/pullrequest/response"
	"revass/internal/service"
	"revass/internal/storage"

	"github.com/go-playground/validator/v10"
)

func Create(log *slog.Logger, prManager service.PRManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var prDTO request.PullRequestCreateRequest
		err := json.NewDecoder(r.Body).Decode(&prDTO)
		if err != nil {
			log.Error("Json decode", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		validate := validator.New()
		err = validate.Struct(prDTO)
		if err != nil {
			handler.MakeBadRequestErrorResponse(w)

			return
		}

		pr, err := prManager.Create(prDTO.PullRequestID, prDTO.PullRequestName, prDTO.AuthorID)
		if err != nil {
			var errEntityExists *storage.ErrEntityExists
			if errors.As(err, &errEntityExists) {
				switch {
				case errors.Is(err, storage.ErrPRExists):
					response.MakePRAlreadyExistsReponse(w, errEntityExists.ID)
					return
				}
			}
			if errors.Is(err, storage.ErrEntityNotFound) {
				handler.MakeNotFoundErrorResponse(w)

				return
			}

			log.Error("PR Create", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		handler.MakeJsonResponse(w, response.PRCreateResponse{PR: pr}, http.StatusCreated)
	}
}

func Merge(log *slog.Logger, prManager service.PRManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var prDTO request.PullRequestMergeRequest
		err := json.NewDecoder(r.Body).Decode(&prDTO)
		if err != nil {
			log.Error("Json decode", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		validate := validator.New()
		err = validate.Struct(prDTO)
		if err != nil {
			handler.MakeBadRequestErrorResponse(w)

			return
		}

		pr, err := prManager.Merge(prDTO.PullRequestID)
		if err != nil {
			if errors.Is(err, storage.ErrEntityNotFound) {
				handler.MakeNotFoundErrorResponse(w)

				return
			}

			log.Error("PR Merge", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		handler.MakeJsonResponse(w, response.PRCreateResponse{PR: pr}, http.StatusOK)
	}
}

func Reassign(log *slog.Logger, prManager service.PRManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var prDTO request.PullRequestReassignRequest
		err := json.NewDecoder(r.Body).Decode(&prDTO)
		if err != nil {
			log.Error("Json decode", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		validate := validator.New()
		err = validate.Struct(prDTO)
		if err != nil {
			handler.MakeBadRequestErrorResponse(w)

			return
		}

		pr, newReviewer, err := prManager.Reassign(prDTO.PullRequestID, prDTO.OldReviewerID)
		if err != nil {
			if errors.Is(err, storage.ErrEntityNotFound) {
				handler.MakeNotFoundErrorResponse(w)

				return
			}

			if errors.Is(err, service.ErrPRMerged) {
				response.MakePRMergedReponse(w)

				return
			}

			log.Error("PR Merge", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		handler.MakeJsonResponse(w, response.PRReassignResponse{PR: pr, ReplacedBy: newReviewer}, http.StatusOK)
	}
}
