package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"watchman/utils"
)

func GetDirAndCmd() (directory string, command string) {
	fmt.Print("Use current directory? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	useCurrentDirectory, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	if strings.TrimSuffix(useCurrentDirectory, "\n") == "n" {
		fmt.Print("Enter directory to watch: ")

		directory, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading directory: ", err)
			return
		}

		directory = strings.TrimSuffix(directory, "\n")
	} else {
		directory = utils.GetCurrentWorkingDirectory()
	}

	fmt.Print("Enter command to run: ")
	command, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading command: ", err)
		return
	}

	command = strings.TrimSuffix(command, "\n")

	return directory, command
}
