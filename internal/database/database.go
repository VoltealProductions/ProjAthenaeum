package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func open(s string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(s), &gorm.Config{})
}

func ConnectToDb() (*gorm.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL"), os.Getenv("DB_TZ"))
	db, err := open(connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
