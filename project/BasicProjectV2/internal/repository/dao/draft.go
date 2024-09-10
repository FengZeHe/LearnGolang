package dao

import (
	"gorm.io/gorm"
)

type GORMDraftDAO struct {
	db *gorm.DB
}

type DraftDAO interface {
	//Insert(ctx context.Context, d Draft) (err error)
}

func NewDraftDAO(db *gorm.DB) DraftDAO {
	return &GORMDraftDAO{
		db: db,
	}
}

//func (dao *GORMDraftDAO) Insert(ctx context.Context, d Draft) (err error) {
//	return err
//}

type Draft struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
	AuthorID   string `json:"authorID"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	DeletedAt  string `json:"deletedAt"`
}
