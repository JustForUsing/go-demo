package model

import (
	"time"
)

type User struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"size:128;uniqueIndex:idx_users_username;not null"`
	Email        string    `gorm:"size:256;uniqueIndex:idx_users_email"`
	Nickname     string    `gorm:"size:128;not null"`
	Logo         string    `gorm:"size:512"`
	PasswordHash string    `gorm:"size:256;not null"`
	FirstLogin   bool      `gorm:"not null"`
	IsAdmin      bool      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (*User) TableName() string {
	return "users"
}
