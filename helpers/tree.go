package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"watchman/utils"
)

func BuildFileTree(directory string) []string {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error: ", err)
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return files
}

func areSlicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func WatchFileTree(directory string) {
	fileTreeHistory := make([][]string, 0)

	utils.RunEvery(5*time.Second, func(t time.Time) {
		fileTreeHistory = append(fileTreeHistory, BuildFileTree(directory))

		if len(fileTreeHistory) > 2 {
			fileTreeHistory = fileTreeHistory[1:]
		}

		if len(fileTreeHistory) > 1 {
			if areSlicesEqual(fileTreeHistory[0], fileTreeHistory[1]) == false {
				fmt.Println("Changed")
			}
		}
	})
}
