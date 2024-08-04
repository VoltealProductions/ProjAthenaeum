package utilities

import (
	"net/http"

	"github.com/a-h/templ"
)

type ctxKey string

var userIdContextKey ctxKey

func RenderView(w http.ResponseWriter, r *http.Request, component templ.Component) error {
	return component.Render(r.Context(), w)
}
