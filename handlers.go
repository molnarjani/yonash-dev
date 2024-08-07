package main

import (
	"log/slog"
	"net/http"

	"github.com/angelofallars/htmx-go"

	"github.com/molnarjani/yonash-dev/templates"
	"github.com/molnarjani/yonash-dev/templates/pages"
)

// indexViewHandler handles a view for the index page.
func indexViewHandler(w http.ResponseWriter, r *http.Request) {
	// Define template meta tags.
	metaTags := pages.MetaTags(
		"gowebly, htmx example page, go with htmx",               // define meta keywords
		"Welcome to example! You're here because it worked out.", // define meta description
	)

	// Define template body content.
	bodyContent := pages.BodyContent(
		"Welcome to example!",                // define h1 text
		"You're here because it worked out.", // define p text
	)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Welcome to example!", // define title text
		metaTags,              // define meta tags
		bodyContent,           // define body content
	)

	// Render index page template.
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, indexTemplate); err != nil {
		// If not, return HTTP 400 error.
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send log message.
	slog.Info("render page", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}

// showContentAPIHandler handles an API endpoint to show content.
func showContentAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(r) {
		// If not, return HTTP 400 error.
		slog.Error("request API", "method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Write HTML content.
	_, err := w.Write([]byte("<p>🎉 Yes, <strong>htmx</strong> is ready to use! (<code>GET /api/hello-world</code>)</p>"))
	if err != nil {
		slog.Error("write response", "error", err)
	}

	// Send htmx response.
	err = htmx.NewResponse().Write(w)
	if err != nil {
		slog.Error("write response", "error", err)
	}

	// Send log message.
	slog.Info("request API", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}
