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

	userRep := repository.NewUserRepository(db)
	teamRep := repository.NewTeamRepository(db)
	prRep := repository.NewPRRepository(db)

	teamManager := service.NewTeamManager(userRep, teamRep)
	prManager := service.NewPRManager(prRep, userRep, teamRep)
	userManager := service.NewUserManager(userRep)

	router.HandleFunc("/team/add", team.Add(log, teamManager)).Methods("POST")
	router.HandleFunc("/team/get", team.Get(log, teamManager)).Methods("GET")

	router.HandleFunc("/users/setIsActive", users.SetIsActive(log, userManager)).Methods("POST")
	router.HandleFunc("/users/getReview", users.GetReview(log, userManager)).Methods("GET")

	router.HandleFunc("/pullRequest/create", pullrequest.Create(log, prManager)).Methods("POST")
	router.HandleFunc("/pullRequest/merge", pullrequest.Merge(log, prManager)).Methods("POST")
	router.HandleFunc("/pullRequest/reassign", pullrequest.Reassign(log, prManager)).Methods("POST")

	return router
}
