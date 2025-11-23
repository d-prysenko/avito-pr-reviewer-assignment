package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ErrorCodeBadRequest          = "BAD_REQUEST"
	ErrorCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrorCodeNotFound            = "NOT_FOUND"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func jsonSerialize(data any) string {
	jsonBytes, _ := json.Marshal(data)

	return string(jsonBytes)
}

func MakeJsonResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%s\n", jsonSerialize(data))
}

func MakeErrorResponse(w http.ResponseWriter, e ErrorResponse, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%s\n", jsonSerialize(
		map[string]any{"error": e}),
	)
}

func MakeInternalServerErrorResponse(w http.ResponseWriter) {
	MakeErrorResponse(w, ErrorResponse{Code: ErrorCodeInternalServerError, Message: "Internal server error"}, http.StatusInternalServerError)
}

func MakeBadRequestErrorResponse(w http.ResponseWriter) {
	MakeErrorResponse(w, ErrorResponse{Code: ErrorCodeBadRequest, Message: "Bad request"}, http.StatusBadRequest)
}

func MakeNotFoundErrorResponse(w http.ResponseWriter) {
	MakeErrorResponse(w, ErrorResponse{Code: ErrorCodeNotFound, Message: "Resource not found"}, http.StatusNotFound)
}
