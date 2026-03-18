package services

import (
	"item-manager-new/internal/repos"
)

type TeamService struct {
	repo *repos.TeamRepos
}

func NewTeamService(repo *repos.TeamRepos) *TeamService {
	return &TeamService{repo: repo}
}
