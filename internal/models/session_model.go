package models

import (
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Session struct {
	UserID uint      `gorm:"index" json:"user_id"`
	Token  string    `gorm:"NOT NULL" json:"token"`
	Expiry time.Time `gorm:"NOT NULL" json:"expiry"`
	gorm.Model
}

func StoreSesson(tk string, id uint, exp time.Time) error {
	sess := Session{
		UserID: id,
		Token:  tk,
		Expiry: exp,
	}

	res := db.Create(&sess)
	return res.Error
}

func LoadSession(tkn string) (Session, bool) {
	session := Session{}
	result := db.Where("token = ?", tkn).First(&session)
	if result.RowsAffected != 1 {
		return Session{}, false
	}
	return session, true
}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

func DeleteSession(tkn string) error {
	session := Session{}
	result := db.Delete(&session, "token = ?", tkn)
	if result.RowsAffected != 1 {
		return result.Error
	}
	return nil
}

func IsLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return false
		}
		return false
	}

	tkn := c.Value
	userSession, exists := LoadSession(tkn)
	if !exists {
		return false
	}

	if userSession.IsExpired() {
		return false
	}

	return true
}

func IsNotLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		return false
	}

	tkn := c.Value
	_, exists := LoadSession(tkn)
	if exists {
		return true
	}

	return false
}
