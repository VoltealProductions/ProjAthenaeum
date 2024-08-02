package handlers

import (
	"fmt"
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/utilities"
	"github.com/VoltealProductions/Athenaeum/internal/views/pages"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, pages.Index())
	if err != nil {
		utilities.RespondWithError(w, 500, fmt.Sprintf("Error rendering template for Index page: %v", err.Error()))
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, pages.About())
	if err != nil {
		utilities.RespondWithError(w, 500, fmt.Sprintf("Error rendering template for About page: %v", err.Error()))
	}
}

func TermsHandler(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, pages.Tos())
	if err != nil {
		utilities.RespondWithError(w, 500, fmt.Sprintf("Error rendering template for Terms page: %v", err.Error()))
	}
}

func FaqHandler(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, pages.Faq())
	if err != nil {
		utilities.RespondWithError(w, 500, fmt.Sprintf("Error rendering template for FaQ page: %v", err.Error()))
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, pages.Contact())
	if err != nil {
		utilities.RespondWithError(w, 500, fmt.Sprintf("Error rendering template for Contact page: %v", err.Error()))
	}
}
