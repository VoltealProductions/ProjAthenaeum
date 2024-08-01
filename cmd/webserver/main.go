package main

import (
	"fmt"

	"github.com/VoltealProductions/Athenaeum/internal/config"
)

func main() {
	config.Set()
	if config.Prod {
		fmt.Println("Environment set to: prod. Hiding errors.")
	} else {
		fmt.Println("Environment set to: dev. Showing errors.")
	}
}
