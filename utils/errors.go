package utils

import (
	"encoding/json"
	"net/http"
	"watchman/schema"
)

func HandleError(w http.ResponseWriter, r *http.Request, statusCode int, message string, err error) {
	if err != nil {
		w.WriteHeader(statusCode)
		response := schema.ResponseType{
			Status:    "ERROR",
			Message:   message + err.Error(),
			RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		}
		SendResponse(w, r, response)
	}
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request, method string) {
	if r.Method != method {
		HandleError(w, r, http.StatusMethodNotAllowed, "Method "+method+" not allowed", nil)
	}
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
