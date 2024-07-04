package utils

import (
	"encoding/json"
	"net/http"
	"watchman/schema"
)

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request, method string) {
	if r.Method != method {
		response := schema.Response_Type{
			Status:    "ERROR",
			Message:   "Method " + method + " not allowed",
			RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
	}
}
