package utils

import (
    "os"
    "fmt"
)

func GetCurrentWorkingDirectory() string {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return directory
}

