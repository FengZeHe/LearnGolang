package service

import "github.com/basicprojectv2/interactive/repository"

type interactiveService struct {
	interactiveRepo repository.InteractiveRepository
}

type InteractiveService interface {
}

func NewInteractiveService(interactiveRepository repository.InteractiveRepository) InteractiveService {
	return &interactiveService{
		interactiveRepo: interactiveRepository,
	}
}
