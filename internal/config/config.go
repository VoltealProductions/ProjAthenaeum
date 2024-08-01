package config

import (
	"flag"

	"github.com/joho/godotenv"
)

var (
	Prod      bool
	Detatched bool
	Port      string
)

func Set() {
	godotenv.Load("./env")

	flag.BoolVar(&Prod, "prod", false, "Production mode; hide all errors.")
	flag.BoolVar(&Detatched, "d", false, "Run Athenaeum in detatched mode.")
	flag.StringVar(&Port, "port", "8080", "The desired port to lsiten on.")
	flag.Parse()
}
