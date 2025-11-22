package router

import (
	"fmt"
	"net/http"
	"revass/internal/handlers/pullrequest"
	"revass/internal/handlers/team"
	"revass/internal/handlers/users"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// router.Use(middleware.CorsMiddleware)

	router.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", request.URL.Path)
	})

	router.HandleFunc("/team/add", team.Add()).Methods("POST", "OPTIONS")
	router.HandleFunc("/team/get", team.Get()).Methods("GET", "OPTIONS")

	router.HandleFunc("/users/setIsActive", users.SetIsActive()).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/getReview", users.GetReview()).Methods("GET", "OPTIONS")

	router.HandleFunc("/pullRequest/create", pullrequest.Create()).Methods("POST", "OPTIONS")
	router.HandleFunc("/pullRequest/merge", pullrequest.Merge()).Methods("POST", "OPTIONS")
	router.HandleFunc("/pullRequest/reassign", pullrequest.Reassign()).Methods("POST", "OPTIONS")

	return router
}
