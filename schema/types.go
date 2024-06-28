package schema

type Response_Type struct {
	Data      interface{} `json:"data,omitempty"`
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
}

type RequestIDKey struct{}

type ConfigType struct {
	Port int `yaml:"port"`
}
