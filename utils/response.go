package utils

import (
	"encoding/json"
	"net/http"
	"watchman/schema"
)

func SendResponse(w http.ResponseWriter, r *http.Request, response schema.ResponseType) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(schema.ResponseType{
			Status:    "ERROR",
			Message:   "Error encoding response",
			RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		})
	}
}
