package service

import (
	"context"

	"github.com/basicprojectv2/interactive/repository"
)

type interactiveService struct {
	interactiveRepo repository.InteractiveRepository
}

type InteractiveService interface {
	AddReadCount(Aid string, ctx context.Context) (err error)
}

func NewInteractiveService(interactiveRepository repository.InteractiveRepository) InteractiveService {
	return &interactiveService{
		interactiveRepo: interactiveRepository,
	}
}

func (i *interactiveService) AddReadCount(Aid string, ctx context.Context) (err error) {
	return i.interactiveRepo.AddReadCount(Aid, ctx)
}
