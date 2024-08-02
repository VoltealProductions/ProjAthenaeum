package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string         `gorm:"not null;unique" json:"username"`
	Email      string         `gorm:"not null;unique" json:"email"`
	Password   string         `gorm:"not null;" json:"password"`
	Public     bool           `gorm:"default:false;" json:"public"`
	Banned     bool           `gorm:"default:false;" json:"banned"`
	VerifiedAt sql.NullTime   `gorm:"default:NULL" json:"verified_at"`
	CreatedAt  sql.NullTime   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  sql.NullTime   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
