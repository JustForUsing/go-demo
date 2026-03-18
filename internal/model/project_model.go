package model

import "time"

type Project struct {
	ID          int64
	TeamID      int64
	Name        string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (*Project) TableName() string {
	return "projects"
}
