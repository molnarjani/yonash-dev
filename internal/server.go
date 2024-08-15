package server

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

//go:embed web/static
var static embed.FS

// runServer runs a new HTTP server with the loaded environment variables.
func RunServer() error {
	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "8080"))
	if err != nil {
		return err
	}

	webStaticSubtree, err := fs.Sub(static, "web/static")
	if err != nil {
		panic(err)
	}
	// Serve static files from the "static" directory.
	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(webStaticSubtree))))

	// Handle API endpoints.
	http.HandleFunc("GET /{page}", pageRoutingHandler)

	// Handle index page view.
	http.HandleFunc("GET /{$}", indexViewHandler)

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
