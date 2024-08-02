package handlers

import (
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/utilities"
	"github.com/VoltealProductions/Athenaeum/internal/views/pages"
)

func GetHomepage(w http.ResponseWriter, r *http.Request) {
	utilities.RenderView(w, r, pages.Index())
}
