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
	HandleLike(aid string, like int, uid string, ctx context.Context) (err error)
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
		return err
	}
	return nil
}

func (i *GORMInteractive) HandleLike(aid string, like int, uid string, ctx context.Context) (err error) {
	var likeArg int
	if like == 0 {
		likeArg = -1
	} else {
		likeArg = 1
	}
	/*
		todo 1. 点赞数+1
			 2. 加一条用户点赞文章的记录
	*/
	now := time.Now().Format("2006-01-02 15:04:05")
	if err = i.db.Model(domain.Interactive{}).Table("interactive").Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]any{
			"like_count": gorm.Expr("like_count + ?", likeArg),
			"utime":      now,
		}),
	}).Create(&domain.Interactive{
		Aid:       aid,
		LikeCount: 1,
		CTime:     now,
		UTime:     now,
	}).Error; err != nil {
		log.Println("interactive like", err)
		return err
	}

	if err = i.db.Model(domain.LikeRecord{}).Table("like_record").Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "aid"}, {Name: "uid"}},
		DoUpdates: clause.Assignments(map[string]any{
			"like":  like,
			"utime": now,
		}),
	}).Create(&domain.LikeRecord{
		Uid:   uid,
		Aid:   aid,
		Like:  like,
		CTime: now,
		UTime: now,
	}).Error; err != nil {
		log.Println("like record error", err)
		return err
	}

	return nil
}
