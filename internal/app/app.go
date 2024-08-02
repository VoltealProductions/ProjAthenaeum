package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/VoltealProductions/Athenaeum/internal/config"
	"github.com/VoltealProductions/Athenaeum/internal/database"
	"github.com/VoltealProductions/Athenaeum/internal/database/models"
	"github.com/VoltealProductions/Athenaeum/internal/database/seeder"
	"github.com/VoltealProductions/Athenaeum/internal/routes"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type App struct {
	router http.Handler
}

type DbConfig struct {
	DB *gorm.DB
}

func New() *App {
	app := &App{
		router: routes.SetRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	if config.Dev {
		err := godotenv.Overload("dev.env")
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := database.ConnectToDb()
	if err != nil {
		logger.LogErr(err.Error(), 503)
	}

	dbCfg := DbConfig{
		DB: db,
	}

	err = dbCfg.migrator()
	if err != nil {
		logger.LogErr(err.Error(), 503)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("WEBSERVER_HOST"), os.Getenv("WEBSERVER_PORT")),
		Handler: a.router,
	}

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func Shutdown() {
}

func (appCfg *DbConfig) migrator() error {
	err := appCfg.DB.AutoMigrate(
		&models.User{},
	)

	if config.Seed {
		seeder.RunSeeders(10)
	}

	return err
}
