package dao

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/basicprojectv2/relationship/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GORMRelationship struct {
	db *gorm.DB
}

var (
	Increase  = 1
	Invariant = 0
	Reduce    = -1
)

type RelationshipDAO interface {
	HandleFollow(followeeUId, followerUId string, action int, ctx context.Context) error
	HandleBlock(blockerUId, blockedUid string, action int, ctx context.Context) error
}

func NewGORMRelationshipDAO(db *gorm.DB) RelationshipDAO {
	return &GORMRelationship{db: db}
}

func (d *GORMRelationship) HandleFollow(followeeUId, followerUId string, action int, ctx context.Context) (err error) {

	// 1. 修改follow关系 2. 更新relationship_record表
	now := time.Now().Format("2006-01-02 15:04:05")

	return d.db.Transaction(func(tx *gorm.DB) error {
		followTarget := domain.User{}
		if err = tx.Model(&domain.User{}).Where("id = ?", followerUId).First(&followTarget).Error; err != nil {
			// 若没查到数据，会返回err "record not found"
			return err
		}
		followStatus := domain.UserFollow{}
		tx.Model(&domain.UserFollow{}).Where("followee_id = ? and follower_id = ?", followeeUId, followerUId).First(&followStatus)
		beforeStatus := followStatus.Status

		if err = tx.Model(&domain.UserFollow{}).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "followee_id"}, {Name: "follower_id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"status":     action,
				"updated_at": now,
			}),
		}).Create(&domain.UserFollow{
			FolloweeId: followeeUId,
			FollowerId: followerUId,
			Status:     strconv.Itoa(action),
			CreatedAt:  now,
			UpdatedAt:  now,
		}).Error; err != nil {
			return err
		}

		switch {
		case beforeStatus == strconv.Itoa(action):
			log.Println("已有记录，不增加了")
			return nil
		case beforeStatus != strconv.Itoa(action) && action == 1:
			if err = tx.Model(&domain.Relationship{}).Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "uid"}},
				DoUpdates: clause.Assignments(map[string]any{
					"followee_num": gorm.Expr("followee_num + ?", Increase),
				}),
			}).Create(&domain.Relationship{
				Uid:         followeeUId,
				FolloweeNum: int64(Increase),
				FollowerNum: 0,
				CreatedAt:   now,
			}).Error; err != nil {
				return err
			}

			if err = tx.Model(&domain.Relationship{}).Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "uid"}},
				DoUpdates: clause.Assignments(map[string]any{
					"follower_num": gorm.Expr("follower_num + ?", Increase),
				}),
			}).Create(&domain.Relationship{
				Uid:         followerUId,
				FolloweeNum: 0,
				FollowerNum: int64(Increase),
				CreatedAt:   now,
			}).Error; err != nil {
				return err
			}

			log.Println("followee 关注+1 , follower 粉丝+1")

		case beforeStatus != strconv.Itoa(action) && action == 0:
			if err = tx.Model(&domain.Relationship{}).Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "uid"}},
				DoUpdates: clause.Assignments(map[string]any{
					"followee_num": gorm.Expr("followee_num + ?", Reduce),
				}),
			}).Create(&domain.Relationship{
				Uid:         followeeUId,
				FolloweeNum: int64(Reduce),
				FollowerNum: 0,
				CreatedAt:   now,
			}).Error; err != nil {
				return err
			}

			if err = tx.Model(&domain.Relationship{}).Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "uid"}},
				DoUpdates: clause.Assignments(map[string]any{
					"follower_num": gorm.Expr("follower_num + ?", Reduce),
				}),
			}).Create(&domain.Relationship{
				Uid:         followerUId,
				FolloweeNum: 0,
				FollowerNum: int64(Reduce),
				CreatedAt:   now,
			}).Error; err != nil {
				return err
			}

			log.Println("followee 关注-1 , follower 粉丝-1")

		}

		return nil

	})
}

func (d *GORMRelationship) HandleBlock(blockerUId, blockedUid string, action int, ctx context.Context) (err error) {

	now := time.Now().Format("2006-01-02 15:04:05")
	return d.db.Transaction(func(tx *gorm.DB) error {
		temp := domain.User{}
		if err = tx.Model(&domain.User{}).Where("id = ?", blockedUid).First(&temp).Error; err != nil {
			return err
		}

		if err = tx.Model(&domain.UserBlock{}).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "blocker_id"}, {Name: "blocked_uid"}},
			DoUpdates: clause.Assignments(map[string]any{
				"status":     strconv.Itoa(action),
				"updated_at": now,
			}),
		}).Create(&domain.UserBlock{
			BlockerId: blockerUId,
			BlockedId: blockedUid,
			Status:    strconv.Itoa(action),
			CreatedAt: now,
		}).Error; err != nil {
			return err
		}
		return nil
	})
}
