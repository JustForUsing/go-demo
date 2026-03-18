package services

import "item-manager-new/internal/repos"

type UserRoleService struct {
	repo *repos.UserRoleRepo
}

func NewUserRoleService(repo *repos.UserRoleRepo) *UserRoleService {
	return &UserRoleService{repo: repo}
}
