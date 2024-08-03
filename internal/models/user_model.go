package models

import (
	"database/sql"

	"github.com/VoltealProductions/Athenaeum/internal/database"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/hash"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string         `gorm:"not null;unique" validate:"required" json:"username"`
	Email      string         `gorm:"not null;unique" validate:"required" json:"email"`
	Password   string         `gorm:"not null;" validate:"required" json:"password"`
	Public     bool           `gorm:"default:false;" json:"public"`
	Banned     bool           `gorm:"default:false;" json:"banned"`
	VerifiedAt sql.NullTime   `gorm:"default:NULL" json:"verified_at"`
	CreatedAt  sql.NullTime   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  sql.NullTime   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func connect() {
	dbConn, err := database.ConnectToDb()
	if err != nil {
		logger.LogFatal(err.Error(), 503)
	}

	db = dbConn
}

func GetUserById() {
}

func GetUserByEmail() {
}

func CreateUser(username, email, password string, public bool) error {
	connect()

	pwd, err := hash.HashPassword(password)
	if err != nil {
		return err
	}

	user := User{
		Username: username,
		Email:    email,
		Password: pwd,
		Public:   public,
	}

	res := db.Create(&user)
	return res.Error
}
