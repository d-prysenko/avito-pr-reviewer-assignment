package team

import (
	"fmt"
	"net/http"
)

func Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	}
}

func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	}
}
