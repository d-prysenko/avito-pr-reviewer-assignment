package users

import (
	"fmt"
	"net/http"
)

func SetIsActive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	}
}

func GetReview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	}
}
