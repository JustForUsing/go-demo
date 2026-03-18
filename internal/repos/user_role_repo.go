package repos

import (
	"errors"
	"gorm.io/gorm"
	"item-manager-new/internal/model"
	"strings"
)

type UserRoleRepo struct {
	db *gorm.DB
}

func NewUserRoleRepo(db *gorm.DB) *UserRoleRepo {
	return &UserRoleRepo{db: db}
}

// Bind 绑定用户角色
func (r *UserRoleRepo) Bind(userID, roleID int64) error {
	record := &model.UserRole{UserID: userID, RoleID: roleID}
	if err := r.db.Create(record).Error; err != nil {
		errMsg := strings.ToLower(err.Error())
		if errors.Is(err, gorm.ErrDuplicatedKey) ||
			strings.Contains(errMsg, "duplicate") ||
			strings.Contains(errMsg, "constraint failed") {
			return nil
		}
		return err
	}
	return nil
}
