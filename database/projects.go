package database

import (
	"fmt"
	"watchman/schema"
)

func (s *service) CreateProject(project schema.Project) error {
	stmt, err := s.db.Prepare("INSERT INTO Projects (ID, Name) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(project.ID, project.Name)
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}

	return nil
}

func (s *service) ListAllProjects() ([]schema.Project, error) {
	rows, err := s.db.Query("SELECT * FROM Projects")
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()
	var projects []schema.Project
	for rows.Next() {
		var project schema.Project
		err := rows.Scan(&project.ID, &project.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		projects = append(projects, project)
	}
	return projects, nil
}

// 8263059922

func (s *service) GetProjectByID(projectID string) (schema.Project, error) {
	row := s.db.QueryRow("SELECT * FROM Projects WHERE ID = ?", projectID)
	var project schema.Project
	err := row.Scan(&project.ID, &project.Name)
	if err != nil {
		return schema.Project{}, fmt.Errorf("error scanning row: %v", err)
	}
	return project, nil
}

func (s *service) UpdateProjectByID(projectID string, project schema.Project) error {
	stmt, err := s.db.Prepare("UPDATE Projects SET Name = ?, ID = ? WHERE ID = ?")
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(project.Name, project.ID, projectID)
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}
	return nil
}

func (s *service) DeleteProjectByID(projectID string) error {
	stmt, err := s.db.Prepare("DELETE FROM Projects WHERE ID = ?")
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(projectID)
	if err != nil {
		return fmt.Errorf("error executing statement: %v", err)
	}
	return nil
}
