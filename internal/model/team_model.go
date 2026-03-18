package model

import "time"

type Team struct {
	ID          int64
	Name        string
	Description string
	LeaderID    *int64
	MemberCount int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (*Team) TableName() string {
	return "teams"
}
