package schema

type Response_Type struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

type RequestIDKey struct{}
