package business

import "errors"

var (
	ErrUserExists        = errors.New("用户已存在")
	ErrUserCreateFail    = errors.New("创建用户失败")
	ErrInvalidCredential = errors.New("用户名/邮箱或密码错误")
)
