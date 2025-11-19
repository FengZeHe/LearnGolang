package repository

import "github.com/basicprojectv2/interactive/repository/dao"

type interactiveRepository struct {
	interactiveDAO dao.InteractiveDAO
}

type InteractiveRepository interface {
}

func NewInteractiveRepository(interactiveDAO dao.InteractiveDAO) InteractiveRepository {
	return &interactiveRepository{
		interactiveDAO: interactiveDAO,
	}
}
