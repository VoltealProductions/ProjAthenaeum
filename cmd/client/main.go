package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/VoltealProductions/Athenaeum/internal/config"
	"github.com/VoltealProductions/Athenaeum/internal/database"
	"github.com/VoltealProductions/Athenaeum/internal/handlers"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Env struct {
	db *gorm.DB
}

func main() {
	config.Set()

	if config.Dev {
		err := godotenv.Overload("dev.env")
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := database.ConnectToDb()
	if err != nil {
		logger.LogFatal(err.Error(), 503)
	}
	env := &Env{db: db}
	r := env.loadRoutes()

	if err := http.ListenAndServe(fmt.Sprintf("%s:%v", os.Getenv("WEBSERVER_HOST"), os.Getenv("WEBSERVER_PORT")), r); err != nil {
		logger.LogFatal(err.Error(), 500)
	}
}

func (env *Env) loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Handle static serving files in the /public folder.
	fs := http.FileServer(http.Dir("public"))
	router.Handle("/public/*", http.StripPrefix("/public/", fs))

	// Base website routes
	router.Get("/", handlers.GetIndexHandler)
	router.Get("/about", handlers.GetAboutHandler)
	router.Get("/terms", handlers.GetTermsHandler)
	router.Get("/faq", handlers.GetFaqHandler)
	router.Get("/contact", handlers.GetContactHandler)

	// System POST Routes only (Register, Login, Logout, activate, etc)
	router.Mount("/s", env.systemRouter(chi.NewRouter()))

	// Archive Routes (Characters, Guilds)
	archiveRouter := chi.NewRouter()
	archiveRouter.Mount("/characters", env.characterRouter(chi.NewRouter()))
	archiveRouter.Mount("/guilds", env.guildRouter(chi.NewRouter()))
	router.Mount("/archive", archiveRouter)

	return router
}

func (env *Env) systemRouter(acr *chi.Mux) *chi.Mux {
	acr.Get("/register", handlers.GetRegisterPage)
	acr.Post("/register", handlers.PostRegisterPage)
	acr.Get("/activate", func(w http.ResponseWriter, r *http.Request) {})
	acr.Get("/login", func(w http.ResponseWriter, r *http.Request) {})
	acr.Get("/reset", func(w http.ResponseWriter, r *http.Request) {})
	acr.Get("/logout", func(w http.ResponseWriter, r *http.Request) {})

	return acr
}

func (env *Env) characterRouter(acr *chi.Mux) *chi.Mux {
	acr.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	return acr
}

func (env *Env) guildRouter(acr *chi.Mux) *chi.Mux {
	acr.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	return acr
}
