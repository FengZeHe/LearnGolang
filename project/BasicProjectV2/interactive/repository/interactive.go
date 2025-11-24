package repository

import (
	"context"

	"github.com/basicprojectv2/interactive/repository/dao"
)

type interactiveRepository struct {
	interactiveDAO dao.InteractiveDAO
}

type InteractiveRepository interface {
	AddReadCount(Aid string, ctx context.Context) (err error)
}

func NewInteractiveRepository(interactiveDAO dao.InteractiveDAO) InteractiveRepository {
	return &interactiveRepository{
		interactiveDAO: interactiveDAO,
	}
}

func (i *interactiveRepository) AddReadCount(Aid string, ctx context.Context) (err error) {
	return i.interactiveDAO.AddReadCount(Aid, ctx)
}
