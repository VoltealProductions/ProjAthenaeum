package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/VoltealProductions/Athenaeum/internal/config"
	"github.com/VoltealProductions/Athenaeum/internal/handlers"
	mid "github.com/VoltealProductions/Athenaeum/internal/middleware"
	"github.com/VoltealProductions/Athenaeum/internal/models"
	"github.com/VoltealProductions/Athenaeum/internal/utilities"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/VoltealProductions/Athenaeum/internal/views/pages/httperrors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config.RunConfig()

	r := loadRoutes()

	if err := http.ListenAndServe(fmt.Sprintf("%s:%v", os.Getenv("WEBSERVER_HOST"), os.Getenv("WEBSERVER_PORT")), r); err != nil {
		logger.LogFatal(err.Error(), 500)
	}
}

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Handle static serving files in the /public folder.
	fs := http.FileServer(http.Dir("public"))
	router.Handle("/public/*", http.StripPrefix("/public/", fs))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utilities.RenderView(w, r, httperrors.NotFoundError(models.IsLoggedIn(r)))
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		utilities.RenderView(w, r, httperrors.NotFoundError(models.IsLoggedIn(r)))
	})

	// Base website routes
	router.Get("/", handlers.GetIndexHandler)
	router.Get("/about", handlers.GetAboutHandler)
	router.Get("/terms", handlers.GetTermsHandler)
	router.Get("/faq", handlers.GetFaqHandler)
	router.Get("/contact", handlers.GetContactHandler)

	// System Routes only (Register, Login, Logout, activate, errors, etc)
	router.Mount("/s", systemRouter(chi.NewRouter()))

	// Archive Routes (Characters, Guilds)
	archiveRouter := chi.NewRouter()
	archiveRouter.Use(mid.AuthedMiddleware)
	archiveRouter.Mount("/characters", characterRouter(chi.NewRouter()))
	archiveRouter.Mount("/guilds", guildRouter(chi.NewRouter()))
	router.Mount("/archive", archiveRouter)

	return router
}

func systemRouter(acr *chi.Mux) *chi.Mux {

	// Account activation routes
	acr.Get("/activate", func(w http.ResponseWriter, r *http.Request) {})
	acr.Post("/activate", func(w http.ResponseWriter, r *http.Request) {})

	// Password Reset
	acr.Get("/reset", func(w http.ResponseWriter, r *http.Request) {})
	acr.Post("/reset", func(w http.ResponseWriter, r *http.Request) {})

	acr.Group(func(r chi.Router) {
		r.Use(mid.GuestMiddleware)
		// Login and Register routes
		r.Get("/register", handlers.GetRegisterPage)
		r.Post("/register", handlers.PostRegisterPage)
		r.Get("/login", handlers.GetLoginPage)
		r.Post("/login", handlers.PostLoginPage)
	})

	acr.Group(func(r chi.Router) {
		r.Use(mid.AuthedMiddleware)
		r.Post("/logout", handlers.Logout)
	})

	return acr
}

func characterRouter(acr *chi.Mux) *chi.Mux {
	acr.Get("/", handlers.CharacterIndex)
	return acr
}

func guildRouter(acr *chi.Mux) *chi.Mux {
	acr.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	return acr
}
