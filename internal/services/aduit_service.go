package services

import (
	"fmt"
	"item-manager-new/internal/model"
	"item-manager-new/internal/repos"
	"strings"
)

type AuditService struct {
	repo *repos.AuditRepos
}

func NewAuditService(repo *repos.AuditRepos) *AuditService {
	return &AuditService{repo: repo}
}

func (s *AuditService) Record(actor *model.User, format string, args ...interface{}) error {
	content := strings.TrimSpace(fmt.Sprintf(format, args...))
	if content == "" {
		return nil
	}

	var userID *int64
	if actor != nil {
		id := actor.ID
		userID = &id
	}

	entry := &model.Audit{
		UserID:  userID,
		Content: content,
	}
	_, err := s.repo.Create(entry)
	return err
}

// List 返回审计日志，仅 admin 可调用。
//func (s *AuditService) List(actor *model.User, filter Filter) (int64, []*model.Audit, error) {
//	if actor == nil || !actor.IsAdmin {
//		return 0, nil, ErrPermissionDenied
//	}
//	list, err := s.repo.List(filter)
//	if err != nil {
//		return 0, nil, err
//	}
//	total, err := s.repo.Count(filter)
//	if err != nil {
//		return 0, nil, err
//	}
//	return total, list, nil
//}
