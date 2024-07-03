package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"watchman/schema"
)

// CREATE NEW LOG: INSERT INTO Logs (Time, Level, Message, Subject, UserID, ProjectID) VALUES (?, ?, ?, ?, ?, ?)
// GET ALL LOGS FROM DB: SELECT * FROM Logs
// GET LOGS BY PROJECT ID: SELECT * FROM Logs WHERE ProjectID = ?
// GET LOGS BY USER ID: SELECT * FROM Logs WHERE UserID = ?
// GET LOGS IN TIME RANGE: SELECT * FROM Logs WHERE Time > ? AND Time < ?
// GET LOGS BY LEVEL: SELECT * FROM Logs WHERE Level = ?
// DELETE LOGS BY PROJECT ID: DELETE FROM Logs WHERE ProjectID = ?
// DELETE LOGS BY USER ID: DELETE FROM Logs WHERE UserID = ?
// DELETE LOGS BY TIME RANGE: DELETE FROM Logs WHERE Time > ? AND Time < ?

func CreateLog(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	var log schema.Log
	log.Time = int32(time.Now().Unix())
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&log)
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

	_, err = stmt.Exec(log.Time, log.Level, log.Message, log.Subject, log.UserID, log.ProjectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error inserting log into database: %v", err)
		return
	}

	response := schema.Response_Type{
		Status:    "OK",
		Message:   "Log created successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

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

func GetAllLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	rows, err := db.Query("SELECT * FROM Logs")
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

func GetLogsByProjectID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	projectID := r.URL.Query().Get("project_id")

	rows, err := db.Query("SELECT * FROM Logs WHERE ProjectID = ?", projectID)
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

func GetLogsByUserID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	userID := r.URL.Query().Get("user_id")

	rows, err := db.Query("SELECT * FROM Logs WHERE UserID = ?", userID)
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

func GetLogsInTimeRange(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	startTime := r.URL.Query().Get("start_time")
	endTime := r.URL.Query().Get("end_time")

	rows, err := db.Query("SELECT * FROM Logs WHERE Time > ? AND Time < ?", startTime, endTime)
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

func GetLogsByLevel(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	level := r.URL.Query().Get("level")

	rows, err := db.Query("SELECT * FROM Logs WHERE Level = ?", level)
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

func DeleteLogsByProjectID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	projectID := r.URL.Query().Get("project_id")

	stmt, err := db.Prepare("DELETE FROM Logs WHERE ProjectID = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error preparing SQL statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(projectID)
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

func DeleteLogsByUserID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	userID := r.URL.Query().Get("user_id")

	stmt, err := db.Prepare("DELETE FROM Logs WHERE UserID = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error preparing SQL statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
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

func DeleteLogsByTimeRange(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	startTime := r.URL.Query().Get("start_time")
	endTime := r.URL.Query().Get("end_time")

	stmt, err := db.Prepare("DELETE FROM Logs WHERE Time > ? AND Time < ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error preparing SQL statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(startTime, endTime)
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
