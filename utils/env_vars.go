package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load_ENV(env_var string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(env_var)
}

func Verify_ENV_Exists(env_var string) bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if os.Getenv(env_var) == "" {
		return false
	}
	return true
}
