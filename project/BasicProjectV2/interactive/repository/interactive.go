package repository

import (
	"context"

	"github.com/basicprojectv2/interactive/domain"
	"github.com/basicprojectv2/interactive/repository/dao"
)

type interactiveRepository struct {
	interactiveDAO dao.InteractiveDAO
}

type InteractiveRepository interface {
	AddReadCount(aid string, ctx context.Context) (err error)
	HandleLike(aid string, like int, uid string, ctx context.Context) (err error)
	HandleCollect(aid string, collect int, uid string, ctx context.Context) (err error)
	GetStatus(aid, uid string, ctx context.Context) (res domain.InteractiveResp, err error)
}

func NewInteractiveRepository(interactiveDAO dao.InteractiveDAO) InteractiveRepository {
	return &interactiveRepository{
		interactiveDAO: interactiveDAO,
	}
}

func (i *interactiveRepository) AddReadCount(aid string, ctx context.Context) (err error) {
	return i.interactiveDAO.AddReadCount(aid, ctx)
}

func (i *interactiveRepository) HandleLike(aid string, like int, uid string, ctx context.Context) (err error) {
	return i.interactiveDAO.HandleLike(aid, like, uid, ctx)
}

func (i *interactiveRepository) HandleCollect(aid string, collect int, uid string, ctx context.Context) (err error) {
	return i.interactiveDAO.HandleCollect(aid, collect, uid, ctx)
}

func (i *interactiveRepository) GetStatus(aid, uid string, ctx context.Context) (res domain.InteractiveResp, err error) {
	return i.interactiveDAO.GetStatus(aid, uid, ctx)
}
