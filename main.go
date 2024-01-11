package main

import (
	"fmt"
	"watchman/cli"
	"watchman/helpers"
)

func main() {
	directory, command := cli.GetDirAndCmd()

	fmt.Println("\nDirectory to watch: ", directory)
	fmt.Println("Command to run: ", command)

	// fileTree := helpers.BuildFileTree(directory)
	//
	// fmt.Println("\nFile tree: ", fileTree)

	helpers.WatchFileTree(directory)
}
