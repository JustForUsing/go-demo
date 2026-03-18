package services

import "item-manager-new/internal/repos"

type ProjectService struct {
	repo *repos.ProjectRepo
}

func NewProjectService(repo *repos.ProjectRepo) *ProjectService {
	return &ProjectService{repo: repo}
}
