package main

import (
	"context"
	"fmt"

	"github.com/VoltealProductions/Athenaeum/internal/app"
	"github.com/VoltealProductions/Athenaeum/internal/config"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
)

func main() {
	config.Set()
	if config.Prod {
		logger.LogInfo("Environment set to: prod")
	} else {
		logger.LogInfo("Environment set to: dev")
	}

	logger.LogInfo(fmt.Sprintf("Athenaeum Webserver now running on port: %v", config.Port))
	app := app.New()
	if err := app.Start(context.TODO()); err != nil {
		logger.LogFatal(err.Error(), 1)
	}
}
