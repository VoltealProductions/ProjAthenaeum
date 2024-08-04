package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

var (
	Prod bool
)

func SetFlags() {
	flag.BoolVar(&Prod, "prod", false, "Production mode; hide all errors.")
	flag.Parse()
}

func LoadEnvVariables() error {
	if os.Getenv("WEBSERVER_HOST") == "" && os.Getenv("WEBSERVER_PORT") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return err
		}
	}
	return nil
}
