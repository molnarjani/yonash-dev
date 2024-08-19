package server

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	gowebly "github.com/gowebly/helpers"
	"github.com/molnarjani/yonash-dev/internal/types"
)

//go:embed web/static
var static embed.FS

// cdnURLMiddleware sets the CDNURL value on the context.
func cdnURLMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the CDNURL value on the context.
		ctx := context.WithValue(r.Context(), types.CDNUrlContextKey, os.Getenv("CDN_URL"))

		// Call the next handler with the updated context.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// runServer runs a new HTTP server with the loaded environment variables.
func RunServer(getenv func(string) string) error {

	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "8080"))
	if err != nil {
		return err
	}

	if getenv("CDN_URL") == "" {
		webStaticSubtree, err := fs.Sub(static, "web/static")
		if err != nil {
			panic(err)
		}
		// Serve static files from the "static" directory.
		http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(webStaticSubtree))))

	}
	// Handle API endpoints.
	http.HandleFunc("GET /{page}", cdnURLMiddleware(pageRoutingHandler))

	// Handle index page view.
	http.HandleFunc("GET /{$}", cdnURLMiddleware(indexViewHandler))

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Send log message.
	slog.Info("Starting server...", slog.Int("port", port), slog.String("CDN_URL", getenv("CDN_URL")))

	return server.ListenAndServe()
}
