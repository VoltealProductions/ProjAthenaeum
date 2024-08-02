package config

import (
	"flag"
)

var (
	Dev bool
)

func Set() {
	flag.BoolVar(&Dev, "dev", false, "Development mode; hide all errors.")

	flag.Parse()
}
