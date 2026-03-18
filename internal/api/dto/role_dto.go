package dto

import "item-manager-new/internal/model"

type RoleDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type int8   `json:"type"`
}

func toRoleDTO(r *model.Role) RoleDTO {
	if r == nil {
		return RoleDTO{}
	}
	return RoleDTO{
		ID:   r.ID,
		Name: r.Name,
		Type: r.Type,
	}
}
