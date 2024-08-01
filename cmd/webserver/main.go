package main

import (
	"context"
	"fmt"
	"log"

	"github.com/VoltealProductions/Athenaeum/internal/app"
	"github.com/VoltealProductions/Athenaeum/internal/config"
)

func main() {
	config.Set()
	if config.Prod {
		fmt.Println("Environment set to: prod. Hiding errors.")
	} else {
		fmt.Println("Environment set to: dev. Showing errors.")
	}

	app := app.New()
	if err := app.Start(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
