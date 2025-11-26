package dao

import (
	"context"
	"errors"
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
	HandleCollect(aid string, collect int, uid string, ctx context.Context) (err error)
	GetStatus(aid, uid string, ctx context.Context) (res domain.InteractiveResp, err error)
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
	now := time.Now().Format("2006-01-02 15:04:05")
	return i.db.Transaction(func(tx *gorm.DB) error {
		// 前置操作：先查询用户之前是否有点赞记录
		var likeRec domain.LikeRecord
		err = tx.Model(domain.LikeRecord{}).Table("like_record").Where("aid = ? AND uid = ?", aid, uid).First(&likeRec).Error
		//lickExists := !errors.Is(err, gorm.ErrRecordNotFound)

		switch {
		// 之前没点赞,现在点赞
		case like == 1 && likeRec.Like == 0:
			if err = tx.Model(domain.Interactive{}).Table("interactive").Where("aid = ?", aid).
				UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
				log.Println("update interactive like count error", err)
				return err
			}
		// 之前点赞，现在取消点赞
		case like == 0 && likeRec.Like == 1:
			if err = tx.Model(domain.Interactive{}).Table("interactive").Where("aid = ?", aid).
				UpdateColumn("like_count", gorm.Expr("like_count + ?", -1)).Error; err != nil {
			}
		default:
			// 无操作
		}

		// 第二段操作 添加/更新like记录表
		if err = tx.Model(domain.LikeRecord{}).Table("like_record").Clauses(clause.OnConflict{
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
	})
}

func (i *GORMInteractive) HandleCollect(aid string, collect int, uid string, ctx context.Context) (err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	// 查询是否有收藏记录
	var rec domain.CollectRecord
	var recColl int

	return i.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(domain.CollectRecord{}).Table("collect_record").Where("aid = ? AND uid = ?", aid, uid).First(&rec).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				recColl = 0
			} else {
				return err
			}
		}
		recColl = rec.Collected

		switch {
		case collect == 1 && recColl == 0:
			//收藏
			if err = tx.Model(domain.Interactive{}).Table("interactive").Where("aid = ?", aid).
				UpdateColumn("collect_count", gorm.Expr("collect_count + ?", 1)).Error; err != nil {
				log.Println("update interactive collect count error", err)
				return err
			}

		case collect == 0 && recColl == 1:
			//取消收藏
			if err = tx.Model(domain.Interactive{}).Table("interactive").Where("aid = ?", aid).
				UpdateColumn("collect_count", gorm.Expr("collect_count - ?", 1)).Error; err != nil {
				log.Println("update interactive collect count error", err)
				return err
			}

		default:
			// 默认 不改动
		}

		if err = tx.Model(domain.CollectRecord{}).Table("collect_record").Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "aid"}, {Name: "uid"}},
			DoUpdates: clause.Assignments(map[string]any{
				"collected": collect,
				"utime":     now,
			}),
		}).Create(&domain.CollectRecord{
			Aid:       aid,
			Uid:       uid,
			Collected: collect,
			Ctime:     now,
			Utime:     now,
		}).Error; err != nil {
			return err
		}
		return nil
	})

}

func (i *GORMInteractive) GetStatus(aid, uid string, ctx context.Context) (res domain.InteractiveResp, err error) {

	var personRes domain.InteractiveStatus
	var interRes domain.Interactive
	i.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(domain.InteractiveStatus{}).Raw(`
	 SELECT a.collected,b.like from webook.collect_record as a LEFT JOIN webook.like_record as b ON 
		a.aid = b.aid where a.aid=? and a.uid=?;`, aid, uid).Scan(&personRes).Error; err != nil {
			return err
		}

		if err = tx.Model(domain.Interactive{}).Table("interactive").Where("aid = ?", aid).Scan(&interRes).Error; err != nil {
			return err
		}
		return nil
	})

	res.Collected = personRes.Collected
	res.Liked = personRes.Liked
	res.ReadCount = interRes.ReadCount
	res.LikeCount = interRes.LikeCount
	res.Collected = personRes.Collected

	return res, nil
}
