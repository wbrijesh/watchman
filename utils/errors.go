package utils

import (
	"net/http"
	"watchman/schema"
)

func HandleError(w http.ResponseWriter, r *http.Request, statusCode int, message string, err error) {
	w.WriteHeader(statusCode)
	response := schema.Response_Type{
		Status:    "ERROR",
		Message:   message + err.Error(),
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	SendResponse(w, r, response)
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request, method string) {
	if r.Method != method {
		HandleError(w, r, http.StatusMethodNotAllowed, "Method "+method+" not allowed", nil)
	}
}
