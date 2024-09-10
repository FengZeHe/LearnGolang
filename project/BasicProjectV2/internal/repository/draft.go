package repository

import "github.com/basicprojectv2/internal/repository/dao"

type draftRepository struct {
	dao dao.DraftDAO
}

type DraftRepository interface {
	GetDraft()
}

func NewDraftRepository(dao dao.DraftDAO) DraftRepository {
	return &draftRepository{dao: dao}
}

func (r *draftRepository) GetDraft() {}
