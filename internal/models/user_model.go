package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/VoltealProductions/Athenaeum/internal/database"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/hash"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/go-playground/validator/v10"
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

var validate *validator.Validate

func initValidation() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func validateStruct(strt User) string {
	initValidation()
	exportStrings := []string{}

	errs := validate.Struct(strt)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			exportStrings = append(exportStrings, fmt.Sprintf("Field: %s has the following errors: %s", err.Field(), err.ActualTag()))
		}

		newExpStr := strings.Join(exportStrings, ";")

		return newExpStr
	}

	return ""
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

func CreateUser(username, email, password string, public bool) string {
	connect()
	initValidation()

	pwd, err := hash.HashPassword(password)
	if err != nil {
		logger.LogFatal(err.Error(), 500)
	}

	user := User{
		Username: username,
		Email:    email,
		Password: pwd,
		Public:   public,
	}

	errs := validateStruct(user)
	if errs != "" {
		return errs
	}

	db.Create(&user)
	return ""
}
