package schema

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Log struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Subject   string `json:"subject"`
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
	Time      int32  `json:"time"`
}
