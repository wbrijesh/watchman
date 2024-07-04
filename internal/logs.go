package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"watchman/schema"
)

func BatchInsertLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}
	var logs []schema.Log
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&logs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding log data: %v", err)
		return
	}
	stmt, err := db.Prepare("INSERT INTO Logs (Time, Level, Message, Subject, UserID, ProjectID) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error preparing SQL statement: %v", err)
		return
	}
	defer stmt.Close()
	for _, log := range logs {
		log.Time = int32(time.Now().Unix())
		_, err = stmt.Exec(log.Time, log.Level, log.Message, log.Subject, log.UserID, log.ProjectID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error inserting log into database: %v", err)
			return
		}
	}
	response := schema.Response_Type{
		Status:    "OK",
		Message:   "Logs created successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}
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
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error querying database: %v", err)
		return
	}
	defer rows.Close()

	var logs []schema.Log
	for rows.Next() {
		var log schema.Log
		err = rows.Scan(&log.Time, &log.Level, &log.Message, &log.Subject, &log.UserID, &log.ProjectID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error scanning row: %v", err)
			return
		}
		logs = append(logs, log)
	}

	w.Header().Set("Content-Type", "application/json")
	response := schema.Response_Type{
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

// DeleteLogs function similar to above GetLogs that takes in the same parameters and deletes the logs from the database
func DeleteLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}
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

	query = query[:len(query)-5]
	stmt, err := db.Prepare(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error preparing SQL statement: %v", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error deleting logs from database: %v", err)
		return
	}
	response := schema.Response_Type{
		Status:    "OK",
		Message:   "Logs deleted successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
