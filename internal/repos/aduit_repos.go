package repos

import (
	"gorm.io/gorm"
	"item-manager-new/internal/model"
)

type AuditRepos struct {
	db *gorm.DB
}

func NewAuditRepos(db *gorm.DB) *AuditRepos {
	return &AuditRepos{db: db}
}

// Create 写入审计记录。
func (r *AuditRepos) Create(audit *model.Audit) (*model.Audit, error) {
	if err := r.db.Create(audit).Error; err != nil {
		return nil, err
	}
	return audit, nil
}
