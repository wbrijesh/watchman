package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"watchman/schema"
	"watchman/utils"
)

func CreateProject(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	utils.HandleMethodNotAllowed(w, r, http.MethodPost)

	var project schema.Project
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&project)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding project data: %v", err)
		return
	}

	if project.ID == "" {
		project.ID = utils.Generate_UUID()
	}

	stmt, err := db.Prepare("INSERT INTO Projects (ID, Name) VALUES (?, ?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error preparing SQL statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(project.ID, project.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error inserting project into database: %v", err)
		return
	}

	response := schema.Response_Type{
		Status:    "OK",
		Message:   "Project created successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetProjectByID(w http.ResponseWriter, r *http.Request, db *sql.DB, projectID string) {
	utils.HandleMethodNotAllowed(w, r, http.MethodGet)

	row := db.QueryRow("SELECT * FROM Projects WHERE ID = ?", projectID)
	var projectByID schema.Project
	err := row.Scan(&projectByID.ID, &projectByID.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error querying database: %v", err)
		return
	}
	response := schema.Response_Type{
		Status:    "OK",
		Message:   "Project retrieved successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data:      projectByID,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ListAllProjects(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	utils.HandleMethodNotAllowed(w, r, http.MethodGet)

	rows, err := db.Query("SELECT * FROM Projects")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error querying database: %v", err)
		return
	}
	defer rows.Close()
	var projects []schema.Project
	for rows.Next() {
		var project schema.Project
		err := rows.Scan(&project.ID, &project.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error scanning row: %v", err)
			return
		}
		projects = append(projects, project)
	}
	response := schema.Response_Type{
		Status:    "OK",
		Message:   "Projects retrieved successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data:      projects,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateProjectByID(w http.ResponseWriter, r *http.Request, db *sql.DB, projectID string) {
	utils.HandleMethodNotAllowed(w, r, http.MethodPut)

	var project schema.Project
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&project)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding project data: %v", err)
		return
	}

	stmt, err := db.Prepare("UPDATE Projects SET Name = ?, ID = ? WHERE ID = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error preparing SQL statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(project.Name, project.ID, projectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error updating project in database: %v", err)
		return
	}

	response := schema.Response_Type{
		Status:    "OK",
		Message:   "Project updated successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteProjectByID(w http.ResponseWriter, r *http.Request, db *sql.DB, projectID string) {
	utils.HandleMethodNotAllowed(w, r, http.MethodDelete)

	stmt, err := db.Prepare("DELETE FROM Projects WHERE ID = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error preparing SQL statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(projectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error deleting project from database: %v", err)
		return
	}

	response := schema.Response_Type{
		Status:    "OK",
		Message:   "Project deleted successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
