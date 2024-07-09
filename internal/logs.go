package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"watchman/schema"
	"watchman/utils"
)

func BatchInsertLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	utils.HandleMethodNotAllowed(w, r, http.MethodPost)

	var logs []schema.Log
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&logs)
	utils.HandleError(w, r, http.StatusBadRequest, "Error decoding JSON: ", err)

	fmt.Println("Logs: ", logs)

	stmt, err := db.Prepare("INSERT INTO Logs (Time, Level, Message, Subject, UserID, ProjectID) VALUES (?, ?, ?, ?, ?, ?)")
	utils.HandleError(w, r, http.StatusInternalServerError, "Error preparing statement: ", err)
	defer stmt.Close()

	for _, log := range logs {
		log.Time = int32(time.Now().Unix())
		_, err = stmt.Exec(log.Time, log.Level, log.Message, log.Subject, log.UserID, log.ProjectID)
		utils.HandleError(w, r, http.StatusInternalServerError, "Error inserting into database: ", err)
	}

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Logs created successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}

func GetLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	utils.HandleMethodNotAllowed(w, r, http.MethodGet)

	projectID := r.URL.Query().Get("project_id")
	userID := r.URL.Query().Get("user_id")
	startTime := r.URL.Query().Get("start_time")
	endTime := r.URL.Query().Get("end_time")
	level := r.URL.Query().Get("level")

	query := "SELECT * FROM Logs WHERE "

	if projectID != "" {
		query += "ProjectID = '" + projectID + "' AND "
	}
	if userID != "" {
		query += "UserID = '" + userID + "' AND "
	}
	if startTime != "" {
		query += "Time > '" + startTime + "' AND "
	}
	if endTime != "" {
		query += "Time < '" + endTime + "' AND "
	}
	if level != "" {
		query += "Level = '" + level + "' AND "
	}

	if projectID != "" || userID != "" || startTime != "" || endTime != "" || level != "" {
		query = query[:len(query)-5]
	} else {
		query = query[:len(query)-7]
	}

	rows, err := db.Query(query)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error querying database: ", err)
	defer rows.Close()

	var logs []schema.Log
	for rows.Next() {
		var log schema.Log
		err = rows.Scan(&log.Time, &log.Level, &log.Message, &log.Subject, &log.UserID, &log.ProjectID)
		utils.HandleError(w, r, http.StatusInternalServerError, "Error scanning row: ", err)
		logs = append(logs, log)
	}

	w.Header().Set("Content-Type", "application/json")
	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Logs retrieved successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data:      logs,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DeleteLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	utils.HandleMethodNotAllowed(w, r, http.MethodDelete)

	projectID := r.URL.Query().Get("project_id")
	userID := r.URL.Query().Get("user_id")
	startTime := r.URL.Query().Get("start_time")
	endTime := r.URL.Query().Get("end_time")
	level := r.URL.Query().Get("level")
	query := "DELETE FROM Logs WHERE "

	if projectID != "" {
		query += "ProjectID = '" + projectID + "' AND "
	}
	if userID != "" {
		query += "UserID = '" + userID + "' AND "
	}
	if startTime != "" {
		query += "Time > '" + startTime + "' AND "
	}
	if endTime != "" {
		query += "Time < '" + endTime + "' AND "
	}
	if level != "" {
		query += "Level = '" + level + "' AND "
	}

	if projectID != "" || userID != "" || startTime != "" || endTime != "" || level != "" {
		query = query[:len(query)-5]
	} else {
		query = query[:len(query)-7]
	}

	stmt, err := db.Prepare(query)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error preparing statement: ", err)
	defer stmt.Close()

	_, err = stmt.Exec()
	utils.HandleError(w, r, http.StatusInternalServerError, "Error executing statement: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Logs deleted successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}
