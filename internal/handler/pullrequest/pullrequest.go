package pullrequest

import (
	"fmt"
	"net/http"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	}
}

func Merge() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	}
}

func Reassign() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	}
}
