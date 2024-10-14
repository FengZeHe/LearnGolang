package repository

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
	"github.com/basicprojectv2/pkg/snowflake"
	"strconv"
	"time"
)

type draftRepository struct {
	draftDAO dao.DraftDAO
}

type DraftRepository interface {
	GetDraft(ctx context.Context, authorID string) ([]domain.Draft, error)
	GetDraftByID(ctx context.Context, draftID, authorID string) (domain.Draft, error)
	AddDraft(ctx context.Context, draft domain.AddDraftReq, authorID string) error
	AddDraftWithPublished(ctx context.Context, req domain.AddDraftReq, authorID string) error
	UpdateDraft(ctx context.Context, draft domain.UpdateDraftReq) error
	UpdateDraftWithPublished(ctx context.Context, req domain.UpdateDraftReq, authorID string) error
	DeleteDraft(ctx context.Context, draft domain.DeleteDraftReq) error
}

func NewDraftRepository(dao dao.DraftDAO) DraftRepository {
	return &draftRepository{draftDAO: dao}
}

func (r *draftRepository) GetDraft(ctx context.Context, authorID string) ([]domain.Draft, error) {
	res, err := r.draftDAO.FindDraftByAuthorID(authorID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *draftRepository) GetDraftByID(ctx context.Context, draftID, authorID string) (domain.Draft, error) {
	res, err := r.draftDAO.FindDraftByID(draftID, authorID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *draftRepository) AddDraft(ctx context.Context, draft domain.AddDraftReq, authorID string) error {
	req := addReqToDomain(draft)
	req.ID = strconv.Itoa(snowflake.GenId())
	req.AuthorID = authorID
	if err := r.draftDAO.InsertDraft(ctx, req); err != nil {
		return err
	}
	return nil
}

func (r *draftRepository) AddDraftWithPublished(ctx context.Context, req domain.AddDraftReq, authorID string) error {
	// todo 组装数据 准备draft和article数据
	user, err := r.draftDAO.FindUserByID(authorID)
	if err != nil {
		return err
	}

	draft := domain.Draft{
		ID:         strconv.Itoa(snowflake.GenId()),
		AuthorID:   authorID,
		AuthorName: user.Nickname,
		Title:      req.Title,
		Content:    req.Content,
		Status:     req.Status,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
	article := domain.Article{
		ID:         strconv.Itoa(snowflake.GenId()),
		AuthorID:   authorID,
		AuthorName: user.Nickname,
		Title:      req.Title,
		Content:    req.Content,
		Status:     req.Status,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := r.draftDAO.InsertDraftAndArticle(ctx, draft, article); err != nil {
		return err
	}
	return nil
}

func (r *draftRepository) UpdateDraft(ctx context.Context, draft domain.UpdateDraftReq) error {
	d := updateReqToDomain(draft)
	err := r.draftDAO.UpdateDraftByAuthorID(d)
	if err != nil {
		return err
	}
	return nil
}

func (r *draftRepository) UpdateDraftWithPublished(ctx context.Context, req domain.UpdateDraftReq, authorID string) (err error) {
	user, err := r.draftDAO.FindUserByID(authorID)
	if err != nil {
		return err
	}

	draft := domain.Draft{
		ID:         req.DraftID,
		AuthorID:   req.AuthorID,
		AuthorName: user.Nickname,
		Title:      req.Title,
		Content:    req.Content,
		Status:     req.Status,
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	article := domain.Article{
		ID:         strconv.Itoa(snowflake.GenId()),
		AuthorID:   req.AuthorID,
		AuthorName: user.Nickname,
		Title:      req.Title,
		Content:    req.Content,
		Status:     req.Status,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	if err = r.draftDAO.UpdateDraftAndArticle(ctx, draft, article); err != nil {
		return err
	}
	return nil

}

func (r *draftRepository) DeleteDraft(ctx context.Context, draft domain.DeleteDraftReq) error {
	if err := r.draftDAO.DeleteDraftByID(draft.DraftID); err != nil {
		return err
	}
	return nil

}

func addReqToDomain(origin domain.AddDraftReq) (target dao.Draft) {
	target.Title = origin.Title
	target.Content = origin.Content
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
