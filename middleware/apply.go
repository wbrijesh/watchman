package middleware

import (
	"net/http"
	"watchman/utils"
)

func ApplyMiddleware(handler http.Handler) http.Handler {
	config := utils.ReadConfig()

	httpHandlerWithMiddleware := CorsMiddleware(
		RequestIDMiddleware(
			Ratelimit(
				config, handler,
			),
		),
	)

	return httpHandlerWithMiddleware
}
