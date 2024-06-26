package utils

import "github.com/google/uuid"

func Generate_UUID() string {
	return uuid.New().String()
}
