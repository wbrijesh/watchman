package internal

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"watchman/database"
	"watchman/middleware"
	"watchman/schema"
	"watchman/utils"
)

type Server struct {
	db   database.DBService
	port int
}

func NewServer() *http.Server {
	NewServer := &Server{
		port: utils.ReadConfig().Port,

		db: database.New(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      middleware.ApplyMiddleware(NewServer.RegisterRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) RegisterRoutes() http.Handler {
	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/health", s.healthHandler)

	multiplexer.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			s.CreateProjectHandler(w, r)
		case http.MethodGet:
			s.ListAllProjectsHandler(w, r)
		default:
			utils.MethodNotAllowedHandler(w, r)
		}
	})

	multiplexer.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.TrimPrefix(r.URL.Path, "/projects/")
		projectID := url

		if projectID == "" {
			utils.HandleError(w, r, http.StatusBadRequest, "Project ID not provided", fmt.Errorf("project ID not provided"))
		}

		switch r.Method {
		case http.MethodGet:
			s.GetProjectByIDHandler(w, r, projectID)
		case http.MethodPut:
			s.UpdateProjectByIDHandler(w, r, projectID)
		case http.MethodDelete:
			s.DeleteProjectByIDHandler(w, r, projectID)
		default:
			utils.MethodNotAllowedHandler(w, r)
		}
	})

	multiplexer.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			s.BatchInsertLogsHandler(w, r)
		case http.MethodGet:
			s.GetLogsHandler(w, r)
		case http.MethodDelete:
			s.DeleteLogsHandler(w, r)
		default:
			utils.MethodNotAllowedHandler(w, r)
		}
	})

	multiplexer.HandleFunc("/login", AdminLogin)

	return multiplexer
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, r, schema.ResponseType{
		Status:    "OK",
		Message:   "Health check successful",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data:      s.db.Health(),
	})
}
