package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/config"
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
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", config.Port),
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
