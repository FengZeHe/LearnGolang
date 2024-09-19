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
	AddArticle(ctx context.Context, req domain.AddDraftReq, authorID string) error
	GetArticles(ctx context.Context, authorID string) ([]domain.Draft, error)
	GetDraft(ctx context.Context, draftID, authorID string) (domain.Draft, error)
	UpdateArticle(ctx context.Context, req domain.UpdateDraftReq) error
	DeleteArticle(ctx context.Context, req domain.DeleteDraftReq) error
}

func NewDraftService(repo repository.DraftRepository) DraftService {
	return &draftService{repo: repo}
}

func (s *draftService) AddArticle(ctx context.Context, req domain.AddDraftReq, authorID string) (err error) {
	// todo 在service层区分status
	switch req.Status {
	case "0": //只保存，不发表
		if err := s.repo.AddDraft(ctx, req, authorID); err != nil {
			return err
		}
		return nil
	case "1": //新建保存并发表
		if err := s.repo.AddDraftWithPublished(ctx, req, authorID); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *draftService) GetArticles(ctx context.Context, authorID string) (d []domain.Draft, err error) {
	d, err = s.repo.GetDraft(ctx, authorID)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *draftService) GetDraft(ctx context.Context, draftID, authorID string) (d domain.Draft, err error) {
	d, err = s.repo.GetDraftByID(ctx, draftID, authorID)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (s *draftService) UpdateArticle(ctx context.Context, req domain.UpdateDraftReq) (err error) {
	// todo 区分是保存还是保存并发表
	switch req.Status {
	case "0":
		if err = s.repo.UpdateDraft(ctx, req); err != nil {
			return err
		}
		return nil
	case "1":
		if err = s.repo.UpdateDraftWithPublished(ctx, req, req.AuthorID); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *draftService) DeleteArticle(ctx context.Context, req domain.DeleteDraftReq) (err error) {
	if err = s.repo.DeleteDraft(ctx, req); err != nil {
		return err
	}
	return nil
}
