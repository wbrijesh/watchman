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

type ProjectKey struct {
	ProjectID string `json:"project_id"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Expiry    int32  `json:"expiry"`
}
