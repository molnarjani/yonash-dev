package content

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/angelofallars/htmx-go"
)

type Page struct {
	Template templ.Component
}

func (p *Page) Render(r *http.Request, w http.ResponseWriter) {
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, p.Template); err != nil {
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
