package config

import (
	"flag"

	"github.com/joho/godotenv"
)

var (
	Prod bool
	Port string
)

func Set() {
	godotenv.Load("./env")

	flag.BoolVar(&Prod, "prod", false, "Production mode; hide all errors.")
	flag.StringVar(&Port, "port", ":8080", "The desired port to lsiten on. Example: -port=\":8080\"")
	flag.Parse()
}
