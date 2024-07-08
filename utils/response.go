package utils

import (
	"encoding/json"
	"net/http"
	"watchman/schema"
)

func SendResponse(w http.ResponseWriter, r *http.Request, response schema.Response_Type) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(schema.Response_Type{
			Status:    "ERROR",
			Message:   "Error encoding response",
			RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		})
	}
}
