package dao

import (
	"gorm.io/gorm"
	"time"
)

type GORMDraftDAO struct {
	db *gorm.DB
}

type DraftDAO interface {
	Insert(d Draft) (err error)
}

func NewDraftDAO(db *gorm.DB) DraftDAO {
	return &GORMDraftDAO{
		db: db,
	}
}

func (dao *GORMDraftDAO) Insert(d Draft) (err error) {
	d.ID = "2"
	d.AuthorName = "2"
	d.AuthorID = "2"
	d.CreatedAt = time.Now().String()
	d.Title = "Title"
	if err = dao.db.Table("draft").Create(&d).Error; err != nil {
		return err
	}
	return nil
}

type Draft struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
	AuthorID   string `json:"authorID"`
	Title      string `json:"title" column:"title"` // title
	Content    string `json:"content"`
	Status     string `json:"status"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	DeletedAt  string `json:"deletedAt"`
}
