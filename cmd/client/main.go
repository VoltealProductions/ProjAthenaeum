package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/VoltealProductions/Athenaeum/internal/app"
	"github.com/VoltealProductions/Athenaeum/internal/config"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/joho/godotenv"
)

func main() {
	config.Set()
	if config.Prod {
		err := godotenv.Overload(".env")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := godotenv.Overload("dev.env")
		if err != nil {
			log.Fatal(err)
		}
	}

	logger.LogInfo(fmt.Sprintf("Athenaeum Webserver now running on port: %v", os.Getenv("WEBSERVER_PORT")))
	app := app.New()
	if err := app.Start(context.TODO()); err != nil {
		logger.LogFatal(err.Error(), 1)
	}
}
