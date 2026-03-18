package repos

import (
	"errors"
	"gorm.io/gorm"
	"item-manager-new/internal/errors/business"
	"item-manager-new/internal/model"
	"strings"
)

type RoleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) CreateRole(role *model.Role) (*model.Role, error) {
	if err := r.db.Create(role).Error; err != nil {
		errMsg := strings.ToLower(err.Error())
		if errors.Is(err, gorm.ErrDuplicatedKey) ||
			strings.Contains(errMsg, "idx_roles_name") ||
			strings.Contains(errMsg, "duplicate") ||
			strings.Contains(errMsg, "unique") {
			return nil, business.ErrUserExists
		}
		return nil, err
	}
	return role, nil
}

// Name 角色名称查询条件
func (r *RoleRepo) Name(roleName string) *QueryBuilder {
	return &QueryBuilder{model: &model.Role{}, db: r.db.Where("name = ?", roleName)}
}

func (r *RoleRepo) ListUserRoles(userID int64) ([]*model.Role, error) {
	var models []model.Role
	if err := r.db.Table((model.Role{}).TableName()).
		Joins("JOIN user_roles ur ON ur.role_id = roles.id AND ur.user_id = ?", userID).
		Order("roles.id ASC").
		Find(&models).Error; err != nil {
		return nil, err
	}
	result := make([]*model.Role, 0, len(models))
	for _, m := range models {
		result = append(result, &m)
	}
	return result, nil
}
