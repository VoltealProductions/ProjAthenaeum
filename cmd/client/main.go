package main

import (
	"context"
	"fmt"
	"log"

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

	logger.LogInfo(fmt.Sprintln("Athenaeum Webserver now running!"))
	app := app.New()
	if err := app.Start(context.TODO()); err != nil {
		logger.LogFatal(err.Error(), 1)
	}
}
