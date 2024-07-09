package internal

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"
	"watchman/schema"
	"watchman/utils"
)

func GenerateKeys() (string, string) {
	numbers := "1234567890"
	smallAlphabets := "abcdefghijklmnopqrstuvwxyz"
	capitalAlphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	accessKeyChars := capitalAlphabets + numbers
	secretKeyChars := capitalAlphabets + smallAlphabets + numbers

	accessKeyCharsLen := len(accessKeyChars)
	secretKeyCharsLen := len(secretKeyChars)

	accessKey := make([]byte, 16)
	secretKey := make([]byte, 32)

	for i := range accessKey {
		accessKey[i] = accessKeyChars[rand.Intn(accessKeyCharsLen)]
	}

	for i := range secretKey {
		secretKey[i] = secretKeyChars[rand.Intn(secretKeyCharsLen)]
	}

	return string(accessKey), string(secretKey)
}

func CreateProjectKey(w http.ResponseWriter, r *http.Request, db *sql.DB, ProjectID string) {
	AccessKey, SecretKey := GenerateKeys()

	var ProjectKey schema.ProjectKey
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ProjectKey)
	utils.HandleError(w, r, http.StatusBadRequest, "Error decoding JSON: ", err)

	stmt, err := db.Prepare("INSERT INTO ProjectKeys (ProjectID, AccessKey, SecretKey, Expiry) VALUES (?, ?, ?, ?)")
	utils.HandleError(w, r, http.StatusInternalServerError, "Error preparing statement: ", err)
	defer stmt.Close()

	_, err = stmt.Exec(ProjectID, AccessKey, SecretKey, ProjectKey.Expiry)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error executing statement: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project key created successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}

func DeleteProjectKey(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var ProjectKey schema.ProjectKey
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ProjectKey)
	utils.HandleError(w, r, http.StatusBadRequest, "Error decoding JSON: ", err)

	stmt, err := db.Prepare("DELETE FROM ProjectKeys WHERE AccessKey = ? AND SecretKey = ?")
	utils.HandleError(w, r, http.StatusInternalServerError, "Error preparing statement: ", err)
	defer stmt.Close()

	_, err = stmt.Exec(ProjectKey.AccessKey, ProjectKey.SecretKey)
	utils.HandleError(w, r, http.StatusInternalServerError, "Error executing statement: ", err)

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Project key deleted successfully",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
	}
	utils.SendResponse(w, r, response)
}
