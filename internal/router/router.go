package router

import (
	"database/sql"
	"log/slog"
	"revass/internal/handler/pullrequest"
	"revass/internal/handler/team"
	"revass/internal/handler/users"
	"revass/internal/service"
	"revass/internal/storage/repository"

	"github.com/gorilla/mux"
)

func New(db *sql.DB, log *slog.Logger) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// router.Use(middleware.CorsMiddleware)

	userRep := repository.NewUserRepository(db)
	teamRep := repository.NewTeamRepository(db)
	prRep := repository.NewPRRepository(db)

	teamManager := service.NewTeamManager(userRep, teamRep)
	prManager := service.NewPRManager(prRep, userRep, teamRep)

	router.HandleFunc("/team/add", team.Add(log, teamManager)).Methods("POST", "OPTIONS")
	router.HandleFunc("/team/get", team.Get(log, teamManager)).Methods("GET", "OPTIONS")

	router.HandleFunc("/users/setIsActive", users.SetIsActive()).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/getReview", users.GetReview()).Methods("GET", "OPTIONS")

	router.HandleFunc("/pullRequest/create", pullrequest.Create(log, prManager)).Methods("POST", "OPTIONS")
	router.HandleFunc("/pullRequest/merge", pullrequest.Merge(log, prManager)).Methods("POST", "OPTIONS")
	router.HandleFunc("/pullRequest/reassign", pullrequest.Reassign(log, prManager)).Methods("POST", "OPTIONS")

	return router
}
