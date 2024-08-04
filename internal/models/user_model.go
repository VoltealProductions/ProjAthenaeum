package models

import (
	"errors"

	"github.com/VoltealProductions/Athenaeum/internal/database"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/hash"
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

var db = database.DB

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
func GetUserForLogin() {
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
