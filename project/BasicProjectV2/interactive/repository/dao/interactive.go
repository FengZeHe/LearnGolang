package dao

import (
	"context"
	"log"
	"time"

	"github.com/basicprojectv2/interactive/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GORMInteractive struct {
	db *gorm.DB
}

type InteractiveDAO interface {
	AddReadCount(aid string, ctx context.Context) (err error)
}

func NewInteractiveDAO(db *gorm.DB) InteractiveDAO {
	return &GORMInteractive{db: db}
}

func (i *GORMInteractive) AddReadCount(aid string, ctx context.Context) (err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	if err = i.db.Model(domain.Interactive{}).Table("interactive").Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]any{
			"read_count": gorm.Expr("read_count + 1"),
			"utime":      now,
		}),
	}).Create(&domain.Interactive{
		Aid:       aid,
		ReadCount: 1,
		CTime:     now,
		UTime:     now,
	}).Error; err != nil {
		log.Println("interactive add count error:", err)
		return err
	}

	return nil
}
