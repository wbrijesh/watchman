package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"watchman/internal"
	"watchman/middleware"
	"watchman/schema"
	"watchman/utils"

	_ "github.com/mattn/go-sqlite3"
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
	config := utils.Read_Config()

	db_connection, err := sql.Open("sqlite3", "./watchman.db")
	if err != nil {
		panic(err)
	}
	defer db_connection.Close()

	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/health", utils.Health_Check_Handler)

	multiplexer.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			internal.CreateProject(w, r, db_connection)
		case http.MethodGet:
			internal.ListAllProjects(w, r, db_connection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	multiplexer.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.TrimPrefix(r.URL.Path, "/projects/")
		projectID := url

		if projectID == "" {
			response := schema.Response_Type{
				Status:    "ERROR",
				Message:   "Project ID not provided",
				RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
				Data:      nil,
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		switch r.Method {
		case http.MethodGet:
			internal.GetProjectByID(w, r, db_connection, projectID)
		case http.MethodPut:
			internal.UpdateProjectByID(w, r, db_connection, projectID)
		case http.MethodDelete:
			internal.DeleteProjectByID(w, r, db_connection, projectID)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	multiplexer.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			internal.BatchInsertLogs(w, r, db_connection)
		case http.MethodGet:
			internal.GetLogs(w, r, db_connection)
		case http.MethodDelete:
			internal.DeleteLogs(w, r, db_connection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	fmt.Println("Starting server on " + strconv.Itoa(config.Port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port),
		middleware.CorsMiddleware(
			middleware.RequestIDMiddleware(
				middleware.Ratelimit(
					multiplexer,
				)))))
}
