package services

import "item-manager-new/internal/repos"

type TeamUserService struct {
	repo *repos.TeamUserRepo
}

func NewTeamUserService(repo *repos.TeamUserRepo) *TeamUserService {
	return &TeamUserService{repo: repo}
}
