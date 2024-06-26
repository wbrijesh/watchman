package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"watchman/internal"
	"watchman/middleware"
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
	port := utils.Load_ENV("PORT")

	db_connection, err := sql.Open("sqlite3", "./watchman.db")
	if err != nil {
		panic(err)
	}
	defer db_connection.Close()

	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/health", utils.Health_Check_Handler)

	multiplexer.HandleFunc("/project", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			internal.GetProjectByID(w, r, db_connection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	multiplexer.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			internal.CreateProject(w, r, db_connection)
		case http.MethodGet:
			internal.ListProjects(w, r, db_connection)
		case http.MethodPut:
			internal.UpdateProject(w, r, db_connection)
		case http.MethodDelete:
			internal.DeleteProject(w, r, db_connection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	multiplexer.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			internal.CreateLog(w, r, db_connection)
		case http.MethodGet:
			internal.GetAllLogs(w, r, db_connection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	multiplexer.HandleFunc("/logs/project", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			internal.GetLogsByProjectID(w, r, db_connection)
		// case http.MethodDelete:
		// 	internal.DeleteLogsByProjectID(w, r, db_c nnection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	multiplexer.HandleFunc("/logs/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			internal.GetLogsByUserID(w, r, db_connection)
		// case http.MethodDelete:
		// 	internal.DeleteLogsByUserID(w, r, db_connection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	multiplexer.HandleFunc("/logs/time", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			internal.GetLogsInTimeRange(w, r, db_connection)
		// case http.MethodDelete:
		// 	internal.DeleteLogsByTimeRange(w, r, db_connection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	multiplexer.HandleFunc("/logs/level", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			internal.GetLogsByLevel(w, r, db_connection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	fmt.Println("Starting server on " + port)
	log.Fatal(http.ListenAndServe(":"+port,
		middleware.RequestIDMiddleware(
			middleware.Ratelimit(
				multiplexer,
			))))
}
