package models

import (
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
