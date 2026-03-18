package repos

import (
	"gorm.io/gorm"
)

type ProjectUserRepo struct {
	db *gorm.DB
}

func NewProjectUserRepo(db *gorm.DB) *ProjectUserRepo {
	return &ProjectUserRepo{db: db}
}
