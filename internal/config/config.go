package config

import (
	"flag"
)

var (
	Dev  bool
	Seed bool
)

func Set() {
	flag.BoolVar(&Dev, "dev", false, "Development mode; hide all errors.")
	flag.BoolVar(&Seed, "seed", false, "Tell the application to seed the Database with test data.")

	flag.Parse()
}
