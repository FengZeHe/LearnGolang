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
	InsertDraft(ctx context.Context, d Draft) (err error)
	InsertArticle(ctx context.Context, a domain.Article) error
	InsertDraftAndArticle(ctx context.Context, d domain.Draft, a domain.Article) error
	UpdateDraftAndArticle(ctx context.Context, d domain.Draft, a domain.Article) error

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

func (dao *GORMDraftDAO) InsertDraft(ctx context.Context, d Draft) (err error) {
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

func (dao *GORMDraftDAO) InsertArticle(ctx context.Context, a domain.Article) (err error) {
	return err
}

func (dao *GORMDraftDAO) InsertDraftAndArticle(ctx context.Context, d domain.Draft, a domain.Article) (err error) {
	// todo 使用Mysql的事务来做 自动事务
	err = dao.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Table("draft").Create(&d).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Table("article").Create(&a).Error; err != nil {
			return err
		}
		return nil
	})

	return err
	// todo Mysql手动事务
	//tx := dao.db.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//if err := tx.Error; err != nil {
	//	return err
	//}
	//if err := tx.Table("draft").Create(&d).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//if err := tx.Table("article").Create(&a).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//return tx.Commit().Error
}

func (dao *GORMDraftDAO) UpdateDraftAndArticle(ctx context.Context, d domain.Draft, a domain.Article) (err error) {

	err = dao.db.Transaction(func(tx *gorm.DB) error {
		// 只有作者自己写的才能改
		if err := tx.WithContext(ctx).Table("draft").Where("id = ? AND author_id = ?", d.ID, d.AuthorID).
			Updates(map[string]interface{}{
				"title":      d.Title,
				"content":    d.Content,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			}).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Table("article").Create(&a).Error; err != nil {
			return err
		}
		return nil
	})
	return err
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
	if err = dao.db.Table("draft").Where("author_id = ? AND status != ?", authorID, "1").Order("created_at desc").Find(&d).Error; err != nil {
		return d, err
	}
	return d, nil
}

func (dao *GORMDraftDAO) UpdateDraftByAuthorID(d Draft) (err error) {
	u := UpdateDraft{
		Title:     d.Title,
		Content:   d.Content,
		Status:    d.Status,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
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
	Title     string `json:"title" column:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updatedAt"`
}
