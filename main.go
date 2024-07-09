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

func main() {
	config := utils.ReadConfig()

	dbConnection, err := sql.Open("sqlite3", "./watchman.db")
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/health", utils.HealthCheckHandler)

	multiplexer.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			internal.CreateProject(w, r, dbConnection)
		case http.MethodGet:
			internal.ListAllProjects(w, r, dbConnection)
		default:
			utils.MethodNotAllowedHandler(w, r)
		}
	})

	multiplexer.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.TrimPrefix(r.URL.Path, "/projects/")
		projectID := url

		if projectID == "" {
			response := schema.ResponseType{
				Status:    "ERROR",
				Message:   "Project ID not provided",
				RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
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
			internal.GetProjectByID(w, r, dbConnection, projectID)
		case http.MethodPut:
			internal.UpdateProjectByID(w, r, dbConnection, projectID)
		case http.MethodDelete:
			internal.DeleteProjectByID(w, r, dbConnection, projectID)
		default:
			utils.MethodNotAllowedHandler(w, r)
		}
	})

	multiplexer.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			internal.BatchInsertLogs(w, r, dbConnection)
		case http.MethodGet:
			internal.GetLogs(w, r, dbConnection)
		case http.MethodDelete:
			internal.DeleteLogs(w, r, dbConnection)
		default:
			utils.MethodNotAllowedHandler(w, r)
		}
	})

	multiplexer.HandleFunc("/login", internal.AdminLogin)

	fmt.Println("Starting server on port " + strconv.Itoa(config.Port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port),
		middleware.CorsMiddleware(
			middleware.RequestIDMiddleware(
				middleware.Ratelimit(
					config,
					multiplexer,
				)))))
}
