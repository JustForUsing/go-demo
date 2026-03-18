package repos

import "gorm.io/gorm"

type TeamUserRepo struct {
	db *gorm.DB
}

func NewTeamUserRepo(db *gorm.DB) *TeamUserRepo {
	return &TeamUserRepo{db: db}
}
