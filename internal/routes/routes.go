package routes

import (
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	staticFileServer(router)
	router.Group(func(r chi.Router) {
		// Base Website Routes
		router.Get("/", handlers.GetIndexHandler)
		router.Get("/about", handlers.GetAboutHandler)
		router.Get("/terms", handlers.GetTermsHandler)
		router.Get("/faq", handlers.GetFaqHandler)
		router.Get("/contact", handlers.GetContactHandler)

		// System POST Routes only (Register, Login, Logout, activate, etc)
		router.Mount("/sys", systemRouter(chi.NewRouter()))

		// Archive Routes (Characters, Guilds)
		archiveRouter := chi.NewRouter()
		archiveRouter.Mount("/characters", characterRouter(chi.NewRouter()))
		archiveRouter.Mount("/guilds", GuildRouter(chi.NewRouter()))
		router.Mount("/archive", archiveRouter)
	})

	return router
}

func staticFileServer(r *chi.Mux) {
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))
}

func systemRouter(acr *chi.Mux) *chi.Mux {
	acr.Post("/register", handlers.GetRegisterPage)
	acr.Post("/activate", func(w http.ResponseWriter, r *http.Request) {})
	acr.Post("/login", func(w http.ResponseWriter, r *http.Request) {})
	acr.Post("/reset", func(w http.ResponseWriter, r *http.Request) {})
	acr.Post("/logout", func(w http.ResponseWriter, r *http.Request) {})

	return acr
}

func characterRouter(acr *chi.Mux) *chi.Mux {
	acr.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	return acr
}

func GuildRouter(acr *chi.Mux) *chi.Mux {
	acr.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	return acr
}
