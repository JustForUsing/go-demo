package services

import (
	"errors"
	"fmt"
	//"item-manager-new/internal/app"
	"item-manager-new/internal/errors/business"
	"item-manager-new/internal/model"
	"item-manager-new/internal/repos"
)

type RoleService struct {
	repo *repos.RoleRepo
}

func NewRoleService(repo *repos.RoleRepo) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) CreateRole(role *model.Role) (*model.Role, error) {
	if role.Type == 0 {
		role.Type = model.RoleTypeCustom
	}
	return s.repo.CreateRole(role)
}

// BindSystemRole 绑定系统角色
// userID 用户ID
func (s *RoleService) BindSystemRole(userID int64) error {
	var role model.Role
	err := s.repo.Name(model.RoleAdmin).Select("id, type").First(&role)
	if err != nil {
		return err
	}
	fmt.Println(role)
	if role.Type != model.RoleTypeSystem {
		return business.ErrBindSystemRole
	}
	// 绑定用户角色
	return repos.Instance.UserRole.Bind(userID, role.ID)
}

// BindNormalRole 绑定普通用户角色
// userID 用户ID
func (s *RoleService) BindNormalRole(userID int64) error {
	var roleId int64
	err := s.repo.Name(model.RoleNormalUser).Value("id", &roleId)
	if err != nil {
		return err
	}
	fmt.Println(roleId)
	// 绑定用户角色
	return repos.Instance.UserRole.Bind(userID, roleId)
}

// PresetAdminRole 预设系统角色
// adminId 管理员ID
func (s *RoleService) PresetAdminRole(adminId int64) {
	// 预设系统角色
	for _, role := range model.SystemRoles {
		// 检查角色是否存在
		exists, err := s.repo.Name(role.Name).Exist()
		if err != nil {
			panic(fmt.Sprintf("检查角色是否存在失败: %v", err))
		}
		if exists {
			continue
		}

		//没有就创建
		_, err = s.CreateRole(&model.Role{
			Name: role.Name,
			Type: model.RoleTypeSystem,
			Desc: role.Desc,
		})
		if err != nil {
			if errors.Is(err, business.ErrDuplicateRole) {
				continue
			}
			panic(fmt.Sprintf("创建角色失败: %v", err))
		}
	}
	// 给管理员绑定管理员角色
	err := s.BindSystemRole(adminId)
	if err != nil {
		panic(fmt.Sprintf("绑定管理员角色失败: %v", err))
	}

	// 给管理员绑定普通用户角色
	err = s.BindNormalRole(adminId)
	if err != nil {
		panic(fmt.Sprintf("绑定普通用户角色失败: %v", err))
	}
}
