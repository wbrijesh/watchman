package internal

import (
	"encoding/json"
	"net/http"
	"watchman/schema"
	"watchman/utils"
)

func (s *Server) BatchInsertLogsHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleMethodNotAllowed(w, r, http.MethodPost)

	var logs []schema.Log
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&logs)
	utils.HandleError(w, r, http.StatusBadRequest, "Error decoding JSON: ", err)

	err = s.db.BatchInsertLogs(logs)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error inserting logs: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Logs created successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}

func (s *Server) GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleMethodNotAllowed(w, r, http.MethodGet)

	projectID := r.URL.Query().Get("project_id")
	userID := r.URL.Query().Get("user_id")
	startTime := r.URL.Query().Get("start_time")
	endTime := r.URL.Query().Get("end_time")
	level := r.URL.Query().Get("level")

	logs, err := s.db.GetLogs(projectID, userID, startTime, endTime, level)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error getting logs: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Logs retrieved successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data:      logs,
	}
	utils.SendResponse(w, r, response)
}

func (s *Server) DeleteLogsHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleMethodNotAllowed(w, r, http.MethodDelete)

	projectID := r.URL.Query().Get("project_id")
	userID := r.URL.Query().Get("user_id")
	startTime := r.URL.Query().Get("start_time")
	endTime := r.URL.Query().Get("end_time")
	level := r.URL.Query().Get("level")

	err := s.db.DeleteLogs(projectID, userID, startTime, endTime, level)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error deleting logs: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Logs deleted successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}
