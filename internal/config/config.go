package config

import (
	"flag"
	"log"
	"os"

	"github.com/VoltealProductions/Athenaeum/internal/database"
	"github.com/VoltealProductions/Athenaeum/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	Prod bool
)

var datB *gorm.DB

func RunConfig() {
	SetFlags()
	err := LoadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}

	datB = db

	MigrateDB()
}

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

func MigrateDB() {
	err := datB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
