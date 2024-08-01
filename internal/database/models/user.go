package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string         `gorm:"not null;unique" json:"username"`
	Email      string         `gorm:"not null;unique" json:"email"`
	Password   string         `gorm:"not null;" json:"-"`
	IsPublic   bool           `gorm:"default:false;" json:"is_public"`
	IsBanned   bool           `gorm:"default:false;" json:"is_banned"`
	VerifiedAt sql.NullTime   `gorm:"default:NULL" json:"verified_at"`
	Deleted    gorm.DeletedAt `gorm:"index" json:"deleted"`
	CreatedAt  sql.NullTime   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  sql.NullTime   `gorm:"autoUpdateTime" json:"updated_at"`
}
