package model

import "time"

type ProjectUserModel struct {
	ProjectID int64     `gorm:"primaryKey"`
	UserID    int64     `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (ProjectUserModel) TableName() string {
	return "project_users"
}
