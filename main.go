package main

import (
	"fmt"
	"watchman/internal"
	"watchman/utils"
)

func main() {
	server := internal.NewServer()

	fmt.Println("Staring server on port ", utils.ReadConfig().Port)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
