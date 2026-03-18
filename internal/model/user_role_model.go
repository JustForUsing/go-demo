package model

// UserRole 映射 user_roles。
type UserRole struct {
	UserID int64 `gorm:"primaryKey"`
	RoleID int64 `gorm:"primaryKey"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
