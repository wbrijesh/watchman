package middleware

import (
	"context"
	"net/http"
	"watchman/schema"

	"github.com/google/uuid"
)

// NEEDS WORK: wrote a struct for RequestIDKey instead of "RequestID" string due to LSP warning to
// not use string key to avoid potential collisions, but a string is used in the middleware function
// for request IDs in echo's middleware package, not sure what is the best way to handle this
// RequestIDKey is moved to schema package

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), schema.RequestIDKey{}, requestID)

		// ctx := context.WithValue(r.Context(), "requestID", requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
