package dto

import (
	"item-manager-new/internal/model"
	"time"
)

type UserDTO struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email,omitempty"`
	Nickname  string    `json:"nickname"`
	Logo      string    `json:"logo"`
	Roles     []RoleDTO `json:"roles"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func ToUserDTO(u *model.User, roles []RoleDTO) *UserDTO {
	return &UserDTO{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Nickname:  u.Nickname,
		Logo:      u.Logo,
		Roles:     roles,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}

func MapRoles(roleEntities []*model.Role) []RoleDTO {
	if len(roleEntities) == 0 {
		return []RoleDTO{}
	}
	result := make([]RoleDTO, 0, len(roleEntities))
	for _, r := range roleEntities {
		result = append(result, toRoleDTO(r))
	}
	return result
}
