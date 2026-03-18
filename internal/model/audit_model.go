package model

import "time"

type Audit struct {
	ID        int64
	UserID    *int64
	Content   string
	CreatedAt time.Time
}

func (*Audit) TableName() string {
	return "audits"
}
