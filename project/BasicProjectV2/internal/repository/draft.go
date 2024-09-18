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
	// todo 多加一个Article的DAO
}

type DraftRepository interface {
	GetDraft(ctx context.Context, authorID string) ([]domain.Draft, error)
	GetDraftByID(ctx context.Context, draftID, authorID string) (domain.Draft, error)
	AddDraft(ctx context.Context, draft domain.AddDraftReq, authorID string) error
	UpdateDraft(ctx context.Context, draft domain.UpdateDraftReq) error
	DeleteDraft(ctx context.Context, draft domain.DeleteDraftReq) error
}

func NewDraftRepository(dao dao.DraftDAO) DraftRepository {
	return &draftRepository{dao: dao}
}

func (r *draftRepository) GetDraft(ctx context.Context, authorID string) ([]domain.Draft, error) {
	res, err := r.dao.FindDraftByAuthorID(authorID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *draftRepository) GetDraftByID(ctx context.Context, draftID, authorID string) (domain.Draft, error) {
	res, err := r.dao.FindDraftByID(draftID, authorID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *draftRepository) AddDraft(ctx context.Context, draft domain.AddDraftReq, authorID string) error {
	req := addReqToDomain(draft)
	req.ID = strconv.Itoa(snowflake.GenId())
	req.AuthorID = authorID
	// todo  在这里使用MySQL事务
	if err := r.dao.Insert(ctx, req); err != nil {
		return err
	}
	return nil
}

func (r *draftRepository) UpdateDraft(ctx context.Context, draft domain.UpdateDraftReq) error {
	d := updateReqToDomain(draft)
	err := r.dao.UpdateDraftByAuthorID(d)
	if err != nil {
		return err
	}
	return nil
}

func (r *draftRepository) DeleteDraft(ctx context.Context, draft domain.DeleteDraftReq) error {
	if err := r.dao.DeleteDraftByID(draft.DraftID); err != nil {
		return err
	}
	return nil

}

func addReqToDomain(origin domain.AddDraftReq) (target dao.Draft) {
	target.Title = origin.Title
	target.Content = origin.Content
	target.AuthorName = origin.AuthorName
	target.Status = origin.Status
	return target
}

func updateReqToDomain(origin domain.UpdateDraftReq) (target dao.Draft) {
	target.Title = origin.Title
	target.Content = origin.Content
	target.AuthorID = origin.AuthorID
	target.ID = origin.DraftID
	target.Status = origin.Status
	return target
}
