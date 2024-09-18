package dao

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"gorm.io/gorm"
	"time"
)

type GORMDraftDAO struct {
	db *gorm.DB
}

type DraftDAO interface {
	Insert(ctx context.Context, d Draft) (err error)
	FindUserByID(id string) (u domain.User, err error)
	FindDraftByAuthorID(authorID string) (d []domain.Draft, err error)
	FindDraftByID(id, authorID string) (d domain.Draft, err error)
	UpdateDraftByAuthorID(d Draft) (err error)
	DeleteDraftByID(id string) (err error)
}

func NewDraftDAO(db *gorm.DB) DraftDAO {
	return &GORMDraftDAO{
		db: db,
	}
}

func (dao *GORMDraftDAO) Insert(ctx context.Context, d Draft) (err error) {
	user, err := dao.FindUserByID(d.AuthorID)
	if err != nil {
		return err
	}
	d.AuthorName = user.Nickname
	d.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err = dao.db.WithContext(ctx).Table("draft").Create(&d).Error; err != nil {
		return err
	}
	return nil
}

func (dao *GORMDraftDAO) FindUserByID(id string) (u domain.User, err error) {
	if err = dao.db.Table("users").Where("id = ?", id).Find(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (dao *GORMDraftDAO) FindDraftByID(id, authorID string) (d domain.Draft, err error) {
	if err = dao.db.Table("draft").Where("id = ? AND author_id = ?", id, authorID).Find(&d).Error; err != nil {
		return d, err
	}
	return d, nil
}

func (dao *GORMDraftDAO) FindDraftByAuthorID(authorID string) (d []domain.Draft, err error) {
	if err = dao.db.Table("draft").Where("author_id = ?", authorID).Order("created_at desc").Find(&d).Error; err != nil {
		return d, err
	}
	return d, nil
}

func (dao *GORMDraftDAO) UpdateDraftByAuthorID(d Draft) (err error) {
	u := UpdateDraft{
		Title:   d.Title,
		Content: d.Content,
		Status:  d.Status,
	}

	if err = dao.db.Table("draft").Where("id = ? AND author_id = ?", d.ID, d.AuthorID).Updates(&u).Error; err != nil {
		return err
	}
	return nil
}

func (dao *GORMDraftDAO) DeleteDraftByID(id string) (err error) {
	if err = dao.db.Table("draft").Where("id = ?", id).Delete(&Draft{}).Error; err != nil {
		return err
	}
	return nil
}

type Draft struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
	AuthorID   string `json:"authorID"` // 对应userid
	Title      string `json:"title" column:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"` // status 0 未发表 1 已发表仅自己可见 2 已发表所有人可见
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	DeletedAt  string `json:"deletedAt"`
}

type UpdateDraft struct {
	Title   string `json:"title" column:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
