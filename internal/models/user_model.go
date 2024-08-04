package models

import (
	"errors"

	"github.com/VoltealProductions/Athenaeum/internal/database"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/hash"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"not null;unique" json:"username"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null;" json:"password"`
	Public   bool   `gorm:"default:false;" json:"public"`
	Banned   bool   `gorm:"default:false;" json:"banned"`
	Verified bool   `gorm:"default:false" json:"verified"`
	gorm.Model
}

var db *gorm.DB

func init() {
	datConn, err := database.ConnectToDb()
	if err != nil {
		logger.LogFatal(err.Error(), 503)
	}

	db = datConn
}

func CreateUser(username, email, password string, public bool) error {
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

func GetXUsers() {
}

func GetUserById() {
}

// Is the only function that fetches the user's password.
func GetUserForLogin(email string) (User, error) {
	user := User{}
	result := db.Where("email = ?", email).First(&user)
	if result.RowsAffected != 1 {
		return User{}, errors.New("the login credentials are incorrect")
	}
	return user, nil
}

func UpdateUser() {
}

func SoftDeleteUser() {
}

func HardDeleteUser() {
}

func UniqueEmail(email string) error {
	user := User{}
	result := db.Where("email = ?", email).First(&user)
	if result.RowsAffected != 0 {
		return errors.New("an account with that email already exists")
	}
	return nil
}

func UniqueUsername(username string) error {
	user := User{}
	result := db.Where("username = ?", username).First(&user)
	if result.RowsAffected != 0 {
		return errors.New("an account with that username already exists")
	}
	return nil
}
