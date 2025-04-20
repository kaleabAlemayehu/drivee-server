package utils

import (
	"encoding/json"
	"net/http"
)

type Success[T any] struct {
	Status  string `json:"status"`
	Message T      `json:"message"`
}

func SendResponse[T any](w http.ResponseWriter, status string, statusCode int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(Success[T]{Status: status, Message: data})
}
