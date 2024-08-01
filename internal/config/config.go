package config

import (
	"flag"
)

var (
	Prod bool
)

func Set() {
	flag.BoolVar(&Prod, "prod", false, "Production mode; hide all errors.")

	flag.Parse()
}
