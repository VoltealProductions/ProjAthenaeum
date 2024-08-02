package main

import (
	"context"
	"log"

	"github.com/VoltealProductions/Athenaeum/internal/app"
	"github.com/VoltealProductions/Athenaeum/internal/config"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/joho/godotenv"
)

func main() {
	config.Set()
	if !config.Dev {
		err := godotenv.Overload("dev.env")
		if err != nil {
			log.Fatal(err)
		}
	}

	logger.LogInfo("Athenaeum Webserver now running")
	app := app.New()
	if err := app.Start(context.TODO()); err != nil {
		logger.LogFatal(err.Error(), 1)
	}
}
