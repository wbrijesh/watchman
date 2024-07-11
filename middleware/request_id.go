package middleware

import (
	"context"
	"net/http"
	"watchman/schema"

	"github.com/google/uuid"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), schema.RequestIDKey{}, requestID)

		// ctx := context.WithValue(r.Context(), "requestID", requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
