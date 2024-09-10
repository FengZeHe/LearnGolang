package service

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
)

type draftService struct {
	repo repository.DraftRepository
}

type DraftService interface {
	AddArticle(ctx context.Context, req domain.AddDraftReq) error
}

func NewDraftService(repo repository.DraftRepository) DraftService {
	return &draftService{repo: repo}
}

func (s *draftService) AddArticle(ctx context.Context, req domain.AddDraftReq) (err error) {
	if err := s.repo.AddDraft(ctx, req); err != nil {
		return err
	}
	return nil
}
