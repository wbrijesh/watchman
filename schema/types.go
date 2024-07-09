package schema

import "github.com/golang-jwt/jwt/v5"

type Response_Type struct {
	Data      interface{} `json:"data,omitempty"`
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
}

type RequestIDKey struct{}

type UsernameKey struct{}

type ConfigType struct {
	Admin struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"admin"`
	JwtKey                     string `yaml:"jwt_key"`
	Port                       int    `yaml:"port"`
	RateLimitRequestsPerSecond int    "yaml:\"rate_limit_req_per_sec\""
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
