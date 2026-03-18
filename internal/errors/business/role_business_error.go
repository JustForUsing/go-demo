package business

import "errors"

var (
	ErrDuplicateRole  = errors.New("角色已存在")
	ErrBindSystemRole = errors.New("只能绑定系统角色")
)
