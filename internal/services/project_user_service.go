package services

import "item-manager-new/internal/repos"

type ProjectUserService struct {
	repo *repos.ProjectUserRepo
}

func NewProjectUserService(repo *repos.ProjectUserRepo) *ProjectUserService {
	return &ProjectUserService{repo: repo}
}
