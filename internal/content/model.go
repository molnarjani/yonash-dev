package content

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/angelofallars/htmx-go"
	"github.com/molnarjani/yonash-dev/internal/web/templates/pages"
	"github.com/molnarjani/yonash-dev/internal/web/templates/pages/projects"
)

type Page struct {
	RoutingKey string
	Template   templ.Component
}

func (p *Page) Render(r *http.Request, w http.ResponseWriter) {
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, p.Template); err != nil {
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

var Pages = []Page{
	{RoutingKey: "projects", Template: pages.ProjectsPage()},
	{RoutingKey: "cv", Template: pages.CVPage()},
	{RoutingKey: "blog", Template: pages.BlogPage()},
	{RoutingKey: "contact", Template: pages.ContactsPage()},
	{RoutingKey: "project-tastingroom", Template: projects.TastingRoomProject()},
	{RoutingKey: "project-esku", Template: projects.EskuProject()},
	{RoutingKey: "project-yonash-dev", Template: projects.YonashDevProject()},
	{RoutingKey: "project-yonash-home", Template: projects.YonashHomeProject()},
}

func GetPageByRoutingKey(key string) (Page, bool) {
	for _, page := range Pages {
		if page.RoutingKey == key {
			return page, true
		}
	}
	return Page{}, false
}
