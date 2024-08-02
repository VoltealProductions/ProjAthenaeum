package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/VoltealProductions/Athenaeum/internal/config"
	"github.com/VoltealProductions/Athenaeum/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	if !config.Prod {
		err := godotenv.Overload("dev.env")
		if err != nil {
			log.Fatal(err)
		}
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("WEBSERVER_HOST"), os.Getenv("WEBSERVER_PORT")),
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func Shutdown() {
}

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	staticFileServer(router)
	router.Group(func(r chi.Router) {
		// Base Website Routes
		router.Get("/", handlers.IndexHandler)
		router.Get("/about", handlers.AboutHandler)
		router.Get("/terms", handlers.TermsHandler)
		router.Get("/faq", handlers.FaqHandler)
		router.Get("/contact", handlers.ContactHandler)

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

func characterRouter(acr *chi.Mux) *chi.Mux {
	acr.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	return acr
}

func GuildRouter(acr *chi.Mux) *chi.Mux {
	acr.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	return acr
}

func systemRouter(acr *chi.Mux) *chi.Mux {
	acr.Post("/register", func(w http.ResponseWriter, r *http.Request) {})
	acr.Post("/activate", func(w http.ResponseWriter, r *http.Request) {})
	acr.Post("/login", func(w http.ResponseWriter, r *http.Request) {})
	acr.Post("/reset", func(w http.ResponseWriter, r *http.Request) {})
	acr.Post("/logout", func(w http.ResponseWriter, r *http.Request) {})

	return acr
}
