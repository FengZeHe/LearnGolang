package service

import (
	"context"

	"github.com/basicprojectv2/interactive/domain"
	"github.com/basicprojectv2/interactive/repository"
)

type interactiveService struct {
	interactiveRepo repository.InteractiveRepository
}

type InteractiveService interface {
	AddReadCount(aid string, ctx context.Context) (err error)
	HandleLike(aid string, like int, uid string, ctx context.Context) (err error)
	HandleCollect(aid string, collect int, uid string, ctx context.Context) (err error)
	GetStatus(aid, uid string, ctx context.Context) (res domain.InteractiveStatus, err error)
}

func NewInteractiveService(interactiveRepository repository.InteractiveRepository) InteractiveService {
	return &interactiveService{
		interactiveRepo: interactiveRepository,
	}
}

func (i *interactiveService) AddReadCount(aid string, ctx context.Context) (err error) {
	return i.interactiveRepo.AddReadCount(aid, ctx)
}

func (i *interactiveService) HandleLike(aid string, like int, uid string, ctx context.Context) (err error) {
	return i.interactiveRepo.HandleLike(aid, like, uid, ctx)
}

func (i *interactiveService) HandleCollect(aid string, collect int, uid string, ctx context.Context) (err error) {
	return i.interactiveRepo.HandleCollect(aid, collect, uid, ctx)
}

func (i *interactiveService) GetStatus(aid, uid string, ctx context.Context) (res domain.InteractiveStatus, err error) {
	return i.interactiveRepo.GetStatus(aid, uid, ctx)
}
