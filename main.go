// projects api endpoints should be:
// 1. GET /projects
// 2. POST /projects
// 3. GET /projects/{id}
// 4. PUT /projects/{id}
// 5. DELETE /projects/{id}
// logs api endpoints should be restful but also factor in query parameters
// 1. GET /logs
// 2 GET /logs?project_id={id}
// 3. GET /logs?user_id={id}
// 4. GET /logs?start_time={timestamp}&end_time={timestamp}
// 5. GET /logs?level={level}
// 6. POST /logs (this is for batch insertion of logs)
// 7. DELETE /logs?project_id={id}
// 8. DELETE /logs?user_id={id}
// 9. DELETE /logs?start_time={timestamp}&end_time={timestamp}
// 10. DELETE /logs?level={level}

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
			internal.CreateLog(w, r, db_connection)
		case http.MethodGet:
			internal.GetAllLogs(w, r, db_connection)
		default:
			utils.Method_Not_Allowed_Handler(w, r)
		}
	})

	multiplexer.HandleFunc("/logs/batch-insert", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			internal.BatchInsertLogs(w, r, db_connection)
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

	fmt.Println("Starting server on " + strconv.Itoa(config.Port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port),
		middleware.CorsMiddleware(
			middleware.RequestIDMiddleware(
				middleware.Ratelimit(
					multiplexer,
				)))))
}
