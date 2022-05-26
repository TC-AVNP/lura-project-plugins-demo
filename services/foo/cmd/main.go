package main

import (
	"fmt"
	"os"

	"poc/services/foo/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	svc := server.NewServer()

	return svc.Run()
}
