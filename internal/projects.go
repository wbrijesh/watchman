package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"watchman/schema"
	"watchman/utils"
)

// CREATE: INSERT INTO Projects (ID, Name) VALUES (?, ?)
// READ: SELECT * FROM Projects
// READ WITH ID: SELECT * FROM Projects WHERE ID = ?
// UPDATE: UPDATE Projects SET Name = ? WHERE ID = ?
// DELETE: DELETE FROM Projects WHERE ID = ?

func CreateProject(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

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

// func ListProjects(w http.ResponseWriter, r *http.Request, db *sql.DB) {
// 	if r.Method != http.MethodGet {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		fmt.Fprintf(w, "Method %s not allowed", r.Method)
// 		return
// 	}
//
// 	projectID := r.URL.Query().Get("id")
//
// 	if projectID == "" {
// 		rows, err := db.Query("SELECT * FROM Projects")
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			fmt.Fprintf(w, "Error querying database: %v", err)
// 			return
// 		}
// 		defer rows.Close()
// 		var projects []schema.Project
// 		for rows.Next() {
// 			var project schema.Project
// 			err := rows.Scan(&project.ID, &project.Name)
// 			if err != nil {
// 				w.WriteHeader(http.StatusInternalServerError)
// 				fmt.Fprintf(w, "Error scanning row: %v", err)
// 				return
// 			}
// 			projects = append(projects, project)
// 		}
// 		response := schema.Response_Type{
// 			Status:    "OK",
// 			Message:   "Projects retrieved successfully",
// 			RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
// 			Data:      projects,
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		err = json.NewEncoder(w).Encode(response)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 	} else {
// 		row := db.QueryRow("SELECT * FROM Projects WHERE ID = ?", projectID)
// 		var projectByID schema.Project
// 		err := row.Scan(&projectByID.ID, &projectByID.Name)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			fmt.Fprintf(w, "Error querying database: %v", err)
// 			return
// 		}
// 		response := schema.Response_Type{
// 			Status:    "OK",
// 			Message:   "Project retrieved successfully",
// 			RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
// 			Data:      projectByID,
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		err = json.NewEncoder(w).Encode(response)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 	}
// }

// internal.GetProjectByID(w, r, db_connection, projectID)

func GetProjectByID(w http.ResponseWriter, r *http.Request, db *sql.DB, projectID string) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}
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
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}
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
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

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
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

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
