package server

import (
	"log/slog"
	"net/http"

	"github.com/angelofallars/htmx-go"

	"github.com/molnarjani/yonash-dev/internal/content"
	"github.com/molnarjani/yonash-dev/internal/web/templates"
	"github.com/molnarjani/yonash-dev/internal/web/templates/pages"
	"github.com/molnarjani/yonash-dev/internal/web/templates/pages/projects"
)

// indexViewHandler handles a view for the index page.
func indexViewHandler(w http.ResponseWriter, r *http.Request) {
	// Define template meta tags.
	metaTags := pages.MetaTags(
		"yonash.dev, developer page, backend development",      // define meta keywords
		"Developer page for my backend development endeavours", // define meta description
	)

	// Define template body content.
	bodyContent := pages.BodyContent()

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
	pages := map[string]content.Page{
		"projects": {Template: pages.ProjectsPage()},
		"cv":       {Template: pages.CVPage()},
		"blog":     {Template: pages.BlogPage()},
		"contact":  {Template: pages.ContactsPage()},
	}

	pageName := r.PathValue("page")
	page := pages[pageName]

	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(r) {
		// If not, return HTTP 400 error.
		slog.Error("request API", "method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	projects := map[string]content.Page{
		"esku":        {Template: projects.EskuProject()},
		"yonash-dev":  {projects.YonashDevProject()},
		"yonash-home": {projects.YonashHomeProject()},
	}

	projectName := r.PathValue("project")
	project := projects[projectName]

	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(r) {
		// If not, return HTTP 400 error.
		slog.Error("request API", "method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
