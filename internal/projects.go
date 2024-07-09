package internal

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"watchman/schema"
	"watchman/utils"

	"github.com/google/uuid"
)

func CreateProject(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	utils.HandleMethodNotAllowed(w, r, http.MethodPost)

	var project schema.Project
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&project)
	utils.HandleError(w, r, http.StatusBadRequest, "Error decoding JSON: ", err)

	if project.ID == "" {
		project.ID = uuid.New().String()
	}

	stmt, err := db.Prepare("INSERT INTO Projects (ID, Name) VALUES (?, ?)")
	utils.HandleError(w, r, http.StatusInternalServerError, "Error preparing statement: ", err)
	defer stmt.Close()

	_, err = stmt.Exec(project.ID, project.Name)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error executing statement: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project created successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}

func GetProjectByID(w http.ResponseWriter, r *http.Request, db *sql.DB, projectID string) {
	utils.HandleMethodNotAllowed(w, r, http.MethodGet)

	row := db.QueryRow("SELECT * FROM Projects WHERE ID = ?", projectID)
	var projectByID schema.Project
	err := row.Scan(&projectByID.ID, &projectByID.Name)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error querying database: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project retrieved successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data:      projectByID,
	}
	utils.SendResponse(w, r, response)
}

func ListAllProjects(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	utils.HandleMethodNotAllowed(w, r, http.MethodGet)

	rows, err := db.Query("SELECT * FROM Projects")
	utils.HandleError(w, r, http.StatusInternalServerError, "Error querying database: ", err)
	defer rows.Close()

	var projects []schema.Project
	for rows.Next() {
		var project schema.Project
		err := rows.Scan(&project.ID, &project.Name)
		utils.HandleError(w, r, http.StatusInternalServerError, "Error scanning row: ", err)
		projects = append(projects, project)
	}
	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Projects retrieved successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data:      projects,
	}
	utils.SendResponse(w, r, response)
}

func UpdateProjectByID(w http.ResponseWriter, r *http.Request, db *sql.DB, projectID string) {
	utils.HandleMethodNotAllowed(w, r, http.MethodPut)

	var project schema.Project
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&project)
	utils.HandleError(w, r, http.StatusBadRequest, "Error decoding JSON: ", err)

	stmt, err := db.Prepare("UPDATE Projects SET Name = ?, ID = ? WHERE ID = ?")
	utils.HandleError(w, r, http.StatusInternalServerError, "Error preparing statement: ", err)
	defer stmt.Close()

	_, err = stmt.Exec(project.Name, project.ID, projectID)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error executing statement: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project updated successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}

func DeleteProjectByID(w http.ResponseWriter, r *http.Request, db *sql.DB, projectID string) {
	utils.HandleMethodNotAllowed(w, r, http.MethodDelete)

	stmt, err := db.Prepare("DELETE FROM Projects WHERE ID = ?")
	utils.HandleError(w, r, http.StatusInternalServerError, "Error preparing statement: ", err)
	defer stmt.Close()

	_, err = stmt.Exec(projectID)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error executing statement: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project deleted successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}
