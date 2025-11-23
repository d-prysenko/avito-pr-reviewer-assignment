package router

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
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

	teamManager := service.NewTeamManager(userRep, teamRep)

	router.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", request.URL.Path)
	})

	router.HandleFunc("/team/add", team.Add(log, teamManager)).Methods("POST", "OPTIONS")
	router.HandleFunc("/team/get", team.Get(log, teamManager)).Methods("GET", "OPTIONS")

	router.HandleFunc("/users/setIsActive", users.SetIsActive()).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/getReview", users.GetReview()).Methods("GET", "OPTIONS")

	router.HandleFunc("/pullRequest/create", pullrequest.Create()).Methods("POST", "OPTIONS")
	router.HandleFunc("/pullRequest/merge", pullrequest.Merge()).Methods("POST", "OPTIONS")
	router.HandleFunc("/pullRequest/reassign", pullrequest.Reassign()).Methods("POST", "OPTIONS")

	return router
}
