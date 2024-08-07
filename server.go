package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	gowebly "github.com/gowebly/helpers"
)

//go:embed all:web/static/*
var static embed.FS

// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "8080"))
	if err != nil {
		return err
	}

	// Only server the subtree under /static route
	webStaticSubtree, err := fs.Sub(static, "web/static")
	if err != nil {
		panic(err)
	}
	// Handle static files from the embed FS (with a custom handler).
	http.Handle("GET /static/", gowebly.StaticFileServerHandler(http.FS(webStaticSubtree)))

	// Handle index page view.
	http.HandleFunc("GET /", indexViewHandler)

	// Handle API endpoints.
	http.HandleFunc("GET /api/hello-world", showContentAPIHandler)

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Send log message.
	slog.Info("Starting server...", "port", port)

	return server.ListenAndServe()
}
