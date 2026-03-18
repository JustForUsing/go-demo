package repos

import "gorm.io/gorm"

type TeamRepos struct {
	db *gorm.DB
}

func NewTeamRepos(db *gorm.DB) *TeamRepos {
	return &TeamRepos{db: db}
}
