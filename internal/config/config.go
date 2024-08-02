package config

import (
	"flag"
)

var (
	Dev bool
)

func Set() {
	flag.BoolVar(&Dev, "dev", true, "Development mode; hide all errors.")

	flag.Parse()
}
