package repository

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
	"github.com/basicprojectv2/pkg/snowflake"
	"strconv"
)

type draftRepository struct {
	dao dao.DraftDAO
}

type DraftRepository interface {
	GetDraft()
	AddDraft(ctx context.Context, draft domain.AddDraftReq, authorID string) error
}

func NewDraftRepository(dao dao.DraftDAO) DraftRepository {
	return &draftRepository{dao: dao}
}

func (r *draftRepository) GetDraft() {}

func (r *draftRepository) AddDraft(ctx context.Context, draft domain.AddDraftReq, authorID string) error {
	req := toDomain(draft)
	req.ID = strconv.Itoa(snowflake.GenId())
	req.AuthorID = authorID
	req.Status = "0"
	if err := r.dao.Insert(ctx, req); err != nil {
		return err
	}
	return nil
}

func toDomain(origin domain.AddDraftReq) (target dao.Draft) {
	target.Title = origin.Title
	target.Content = origin.Content
	target.AuthorName = origin.AuthorName
	return target
}
