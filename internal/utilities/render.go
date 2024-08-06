package utilities

import (
	"net/http"

	"github.com/a-h/templ"
)

func RenderView(w http.ResponseWriter, r *http.Request, component templ.Component) error {
	return component.Render(r.Context(), w)
}
