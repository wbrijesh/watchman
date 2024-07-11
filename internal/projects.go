package internal

import (
	"encoding/json"
	"net/http"
	"watchman/schema"
	"watchman/utils"

	"github.com/google/uuid"
)

func (s *Server) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleMethodNotAllowed(w, r, http.MethodPost)

	var project schema.Project
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&project)
	utils.HandleError(w, r, http.StatusBadRequest, "Error decoding JSON: ", err)

	if project.ID == "" {
		project.ID = uuid.New().String()
	}

	err = s.db.CreateProject(project)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error creating project: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project created successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}

func (s *Server) GetProjectByIDHandler(w http.ResponseWriter, r *http.Request, projectID string) {
	utils.HandleMethodNotAllowed(w, r, http.MethodGet)

	project, err := s.db.GetProjectByID(projectID)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error retrieving project: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project retrieved successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data:      project,
	}
	utils.SendResponse(w, r, response)
}

func (s *Server) ListAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleMethodNotAllowed(w, r, http.MethodGet)

	projects, err := s.db.ListAllProjects()
	utils.HandleError(w, r, http.StatusInternalServerError, "Error retreiving projects: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Projects retrieved successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data:      projects,
	}
	utils.SendResponse(w, r, response)
}

func (s *Server) UpdateProjectByIDHandler(w http.ResponseWriter, r *http.Request, projectID string) {
	utils.HandleMethodNotAllowed(w, r, http.MethodPut)

	var project schema.Project
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&project)
	utils.HandleError(w, r, http.StatusBadRequest, "Error decoding JSON: ", err)

	err = s.db.UpdateProjectByID(projectID, project)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error updating project: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project updated successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}

func (s *Server) DeleteProjectByIDHandler(w http.ResponseWriter, r *http.Request, projectID string) {
	utils.HandleMethodNotAllowed(w, r, http.MethodDelete)

	err := s.db.DeleteProjectByID(projectID)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error deleting project: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project deleted successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}
