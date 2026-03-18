package services

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"item-manager-new/internal/errors/business"
	"item-manager-new/internal/model"
	"item-manager-new/internal/pkg/global"
	"item-manager-new/internal/pkg/utils"
	"item-manager-new/internal/repos"
	"strings"
)

type UserService struct {
	repo *repos.UserRepo
}

func NewUserService(repo *repos.UserRepo) *UserService {
	return &UserService{repo: repo}
}

// CreateUser 创建用户
func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	// 检查用户是否存在
	exists, err := u.repo.Username(user.Username).Exist()
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, business.ErrUserExists
	}
	// 构建验证用户数据
	passwordHash, err := utils.HashString(user.PasswordHash)
	if err != nil {
		return nil, business.ErrHashStringFailed
	}
	user.PasswordHash = passwordHash

	email := strings.TrimSpace(user.Email)
	user.Email = email

	if newUser, err := u.repo.Create(user); err != nil {
		return nil, business.ErrUserCreateFail
	} else {
		return newUser, nil
	}
}

// PresetAdmin 预设管理员用户
func (u *UserService) PresetAdmin() (int64, error) {
	adminCfg := global.GetAdminConfig()
	newUser, err := u.CreateUser(&model.User{
		Username:     adminCfg.Username,
		PasswordHash: adminCfg.Password,
		Email:        adminCfg.Email,
		Nickname:     adminCfg.Nickname,
		FirstLogin:   true,
		IsAdmin:      true,
	})
	//判断错误是否为用户已存在类型的错误，如果是的话，直接返回 nil，如果不是直接panic
	if err != nil {
		if errors.Is(err, business.ErrUserExists) {
			var adminId int64
			err = u.repo.Username(adminCfg.Username).Value("id", &adminId)
			if err != nil {
				panic(fmt.Sprintf("获取管理员用户ID失败: %v", err))
			}
			return adminId, nil
		}
		panic(fmt.Sprintf("创建管理员用户失败: %v", err))
	}
	return newUser.ID, nil
}

// Auth 用户认证
func (u *UserService) Auth(username, email, password string) (*model.User, error) {
	var (
		user model.User
		err  error
	)
	// 检查用户是否存在
	if email != "" {
		err = u.repo.Email(email).First(&user)
	} else {
		err = u.repo.Username(username).First(&user)
	}
	if err != nil {
		return nil, business.ErrInvalidCredential
	}

	// 检查密码是否正确
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return nil, business.ErrInvalidCredential
	}
	return &user, nil
}
