package team

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"revass/internal/handler"
	"revass/internal/handler/team/request"
	"revass/internal/handler/team/response"
	"revass/internal/service"
	"revass/internal/storage"

	"github.com/go-playground/validator/v10"
)

func Add(log *slog.Logger, teamManager service.TeamManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var teamDTO request.TeamAddRequest
		err := json.NewDecoder(r.Body).Decode(&teamDTO)
		if err != nil {
			log.Error("Json decode", "err", err.Error())
			handler.MakeBadRequestErrorResponse(w)

			return
		}

		validate := validator.New()
		err = validate.Struct(teamDTO)
		if err != nil {
			handler.MakeBadRequestErrorResponse(w)

			return
		}

		team := teamDTO.ToTeamModel()
		err = teamManager.AddTeamWithMembers(team)
		if err != nil {
			var errEntityExists *storage.ErrEntityExists
			if errors.As(err, &errEntityExists) {
				switch {
				case errors.Is(err, storage.ErrUserExists):
					response.MakeUserAlreadyExistsReponse(w, errEntityExists.ID)
					return

				case errors.Is(err, storage.ErrTeamExists):
					response.MakeTeamAlreadyExistsReponse(w, errEntityExists.ID)
					return
				}
			}

			log.Error("Team add", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		handler.MakeJsonResponse(w, response.TeamAddResponse{Team: team}, http.StatusCreated)
	}
}

func Get(log *slog.Logger, teamManager service.TeamManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teamName := r.URL.Query().Get("team_name")
		if teamName == "" {
			handler.MakeBadRequestErrorResponse(w)

			return
		}

		team, err := teamManager.GetTeam(teamName)
		if err != nil {
			if errors.Is(err, storage.ErrEntityNotFound) {
				handler.MakeNotFoundErrorResponse(w)

				return
			}

			log.Error("Team get", "err", err.Error())
			handler.MakeInternalServerErrorResponse(w)

			return
		}

		handler.MakeJsonResponse(w, team, http.StatusOK)
	}
}
