package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"revass/internal/handler"
	"revass/internal/handler/users/response"
	"revass/internal/handler/users/request"
	"revass/internal/service"
	"revass/internal/storage"

	"github.com/go-playground/validator/v10"
)

func SetIsActive(log *slog.Logger, userManager service.UserManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userDTO request.SetIsActiveRequest
		err := json.NewDecoder(r.Body).Decode(&userDTO)
		if err != nil {
			log.Error("Json decode", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		validate := validator.New()
		err = validate.Struct(userDTO)
		if err != nil {
			handler.MakeBadRequestErrorResponse(w)

			return
		}

		user, err := userManager.SetIsActive(userDTO.UserID, userDTO.IsActive)
		if err != nil {
			if errors.Is(err, storage.ErrEntityNotFound) {
				handler.MakeNotFoundErrorResponse(w)

				return
			}

			log.Error("PR Merge", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		handler.MakeJsonResponse(w, response.UserResponse{User: user}, http.StatusOK)
	}
}

func GetReview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	}
}
