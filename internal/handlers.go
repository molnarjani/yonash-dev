package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/angelofallars/htmx-go"

	"github.com/molnarjani/yonash-dev/internal/content"
	"github.com/molnarjani/yonash-dev/internal/web/templates"
	"github.com/molnarjani/yonash-dev/internal/web/templates/pages"
)

func HTMXOrRedirect(w http.ResponseWriter, r *http.Request, page content.Page) {
	// Check if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(r) {
		// If not, redirect to the root page.
		var redirectRoute string
		if page, ok := content.GetPageByRoutingKey(page.RoutingKey); ok {
			redirectRoute = fmt.Sprintf("/?page=%s", page.RoutingKey)
		} else {
			redirectRoute = "/"
			slog.Error("redirect error, page not found", slog.Any("page", page))
		}
		http.Redirect(w, r, redirectRoute, http.StatusTemporaryRedirect)
	}
}

// indexViewHandler handles a view for the index page.
func indexViewHandler(w http.ResponseWriter, r *http.Request) {

	// Define template meta tags.
	metaTags := pages.MetaTags(
		"yonash.dev, developer page, backend development",      // define meta keywords
		"Developer page for my backend development endeavours", // define meta description
	)

	// Get the value of the 'page' query parameter from the request.
	pageParam := r.URL.Query().Get("page")

	// Get the page by the routing key.
	var bodyContent templ.Component
	if page, ok := content.GetPageByRoutingKey(pageParam); ok {
		// Define template body content.
		bodyContent = pages.BodyContent(
			page.Template,
		)
	} else {
		// Define template body content.
		bodyContent = pages.BodyContent(
			pages.IndexPageContent(),
		)
	}

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"yonash.dev",
		metaTags,
		bodyContent,
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

// pageRoutingHandler handles serving Page contents from API
func pageRoutingHandler(w http.ResponseWriter, r *http.Request) {
	pageName := r.PathValue("page")
	page, _ := content.GetPageByRoutingKey(pageName)

	// Redirect if not HTMX
	HTMXOrRedirect(w, r, page)

	// Write HTML content.
	page.Render(r, w)

	// Send htmx response.
	err := htmx.NewResponse().Write(w)
	if err != nil {
		slog.Error("write response", "error", err)
	}

	// Send log message.
	slog.Info("request API", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}

// pageRoutingHandler handles serving Page contents from API
func projectRoutingHandler(w http.ResponseWriter, r *http.Request) {
	projectName := r.PathValue("project")
	projectNamePrefixed := fmt.Sprintf("project-%s", projectName)
	project, ok := content.GetPageByRoutingKey(projectNamePrefixed)
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Redirect if not HTMX
	HTMXOrRedirect(w, r, project)

	// Write HTML content.
	project.Render(r, w)

	// Send htmx response.
	err := htmx.NewResponse().Write(w)
	if err != nil {
		slog.Error("write response", "error", err)
	}

	// Send log message.
	slog.Info("request API", "method", r.Method, "status", http.StatusOK, "path", r.URL.Path)
}
