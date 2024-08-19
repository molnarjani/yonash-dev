package main

import (
	"log/slog"
	"os"

	server "github.com/molnarjani/yonash-dev/internal"
)

func main() {
	// Run your server.
	if err := server.RunServer(os.Getenv); err != nil {
		slog.Error("Failed to start server!", "details", err.Error())
		os.Exit(1)
	}
}
