package main

import (
	"fmt"
	"os"

	"poc/services/lura/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	err := server.NewLuraInstance().Start()
	if err != nil {
		return err
	}

	return nil
}
