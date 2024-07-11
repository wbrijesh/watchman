package database

import (
	"fmt"
	"time"
	"watchman/schema"
)

func (s *service) BatchInsertLogs(logs []schema.Log) error {
	stmt, err := s.db.Prepare("INSERT INTO Logs (Time, Level, Message, Subject, UserID, ProjectID) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()
	for _, log := range logs {
		log.Time = int32(time.Now().Unix())
		_, err = stmt.Exec(log.Time, log.Level, log.Message, log.Subject, log.UserID, log.ProjectID)
		if err != nil {
			return fmt.Errorf("error inserting into database: %v", err)
		}
	}
	return nil
}

func (s *service) GetLogs(projectID string, userID string, startTime string, endTime string, level string) ([]schema.Log, error) {
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
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()
	var logs []schema.Log
	for rows.Next() {
		var log schema.Log
		err := rows.Scan(&log.Time, &log.Level, &log.Message, &log.Subject, &log.UserID, &log.ProjectID)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func (s *service) DeleteLogs(projectID string, userID string, startTime string, endTime string, level string) error {
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
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}
	return nil
}
