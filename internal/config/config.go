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
	setFlags()
	err := loadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}

	datB = db

	migrateDB()
}

func setFlags() {
	flag.BoolVar(&Prod, "prod", false, "Production mode; hide all errors.")
	flag.Parse()
}

func loadEnvVariables() error {
	if os.Getenv("WEBSERVER_HOST") == "" && os.Getenv("WEBSERVER_PORT") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return err
		}
	}
	return nil
}

func migrateDB() {
	err := datB.AutoMigrate(&models.User{}, models.Session{})
	if err != nil {
		log.Fatal(err)
	}
}
