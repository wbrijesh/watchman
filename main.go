package main

import (
	"fmt"
	"log"
	"net/http"
	"watchman/middleware"
	"watchman/utils"
)

const (
	RED          = "\033[31m"
	RESET_COLOUR = "\033[0m"
)

func init() {
	env_vars := []string{"PORT"}

	for _, env_var := range env_vars {
		if !utils.Verify_ENV_Exists(env_var) {
			log.Fatal(RED + "Error: " + env_var + " not found in .env file" + RESET_COLOUR)
		}
	}
}

func main() {
	port := utils.Load_ENV("PORT")

	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/", utils.Example_Handler)
	multiplexer.HandleFunc("/health", utils.Health_Check_Handler)

	fmt.Println("Starting server on " + port)
	log.Fatal(http.ListenAndServe(":"+port,
		middleware.RequestIDMiddleware(
			middleware.Ratelimit(
				multiplexer,
			))))
}
