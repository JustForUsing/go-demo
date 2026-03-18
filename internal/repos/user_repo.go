package repos

import (
	"gorm.io/gorm"
	"item-manager-new/internal/model"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create 创建用户
func (u *UserRepo) Create(user *model.User) (*model.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// FindById 根据ID查询用户
func (u *UserRepo) FindById(id int64) (*model.User, error) {
	var user model.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

//  链式调用，存在轻微的内存分配开销，少量使用，可以接受

// Username 用户名查询构建器
func (u *UserRepo) Username(username string) *QueryBuilder {
	return &QueryBuilder{model: &model.User{}, db: u.db.Where("username = ?", username)}
}

// Email 邮箱查询构建器
func (u *UserRepo) Email(email string) *QueryBuilder {
	return &QueryBuilder{model: &model.User{}, db: u.db.Where("email = ?", email)}
}

// Nickname 昵称查询构建器
func (u *UserRepo) Nickname(nickname string) *QueryBuilder {
	return &QueryBuilder{model: &model.User{}, db: u.db.Where("nickname = ?", nickname)}
}

// Scope ,GORM 官方推荐的方式，用于构建查询条件

// UsernameScope 用户名查询条件
func (u *UserRepo) UsernameScope(username string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ?", username)
	}
}

// NicknameScope 昵称查询条件
func (u *UserRepo) NicknameScope(nickname string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("nickname = ?", nickname)
	}
}

// FindIdByScopes 根据多个条件查询用户ID
func (u *UserRepo) FindIdByScopes(scopes ...func(db *gorm.DB) *gorm.DB) (int64, error) {
	var id int64
	if err := u.db.Model(&model.User{}).
		Select("id").
		Scopes(scopes...).
		Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

// JungleExistByUsername 检查用户是否存在
func (u *UserRepo) JungleExistByUsername(username string) (bool, error) {
	var exists bool
	if err := u.db.Model(&model.User{}).
		Select("EXISTS(SELECT 1 FROM users WHERE username = ?)", username).
		Scan(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}
