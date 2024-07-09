package utils

import (
	"encoding/json"
	"net/http"
	"watchman/schema"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	response := schema.ResponseType{
		Status:    "ERROR",
		Message:   "Method Not Allowed",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
