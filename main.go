package main

import (
	"watchman/cli"
	"watchman/helpers"
)

func main() {
	directory, command := cli.GetDirAndCmd()
	helpers.WatchFileTree(directory, command)
}
