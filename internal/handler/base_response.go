package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ERR_CODE_BAD_REQUEST           = "BAD_REQUEST"
	ERR_CODE_INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	ERR_CODE_NOT_FOUND             = "NOT_FOUND"
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
	MakeErrorResponse(w, ErrorResponse{Code: ERR_CODE_INTERNAL_SERVER_ERROR, Message: "Internal server error"}, http.StatusInternalServerError)
}

func MakeBadRequestErrorResponse(w http.ResponseWriter) {
	MakeErrorResponse(w, ErrorResponse{Code: ERR_CODE_BAD_REQUEST, Message: "Bad request"}, http.StatusBadRequest)
}

func MakeNotFoundErrorResponse(w http.ResponseWriter) {
	MakeErrorResponse(w, ErrorResponse{Code: ERR_CODE_NOT_FOUND, Message: "Resource not found"}, http.StatusNotFound)
}
