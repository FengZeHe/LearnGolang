package service

import (
	"context"
	"github.com/basicprojectv2/internal/repository"
)

type draftService struct {
	repo repository.DraftRepository
}

type DraftService interface {
	AddArticle(ctx context.Context)
}

func NewDraftService(repo repository.DraftRepository) DraftService {
	return &draftService{repo: repo}
}

func (s *draftService) AddArticle(ctx context.Context) {

}
