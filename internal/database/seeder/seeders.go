package seeder

import (
	"github.com/VoltealProductions/Athenaeum/internal/database"
	"github.com/VoltealProductions/Athenaeum/internal/database/models"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func RunSeeders(records int) {
	db, err := database.ConnectToDb()
	if err != nil {
		logger.LogFatal(err.Error(), 503)
	}

	seedUsers(db, records)
}

func seedUsers(db *gorm.DB, rc int) {

	var dbCount int64
	db.Model(&models.User{}).Where("ID = ?", 1).Count(&dbCount)
	if dbCount == 1 {
		return
	}

	for i := 0; i < rc; i++ {
		user := models.User{
			Username: faker.Username(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}

		result := db.Create(&user)
		if result.Error != nil {
			logger.LogFatal(result.Error.Error(), 503)
		}
	}
}
