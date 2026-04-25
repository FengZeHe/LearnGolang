package dao

import (
	"context"
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
}

func NewGORMRelationshipDAO(db *gorm.DB) RelationshipDAO {
	return &GORMRelationship{db: db}
}

func (d *GORMRelationship) HandleFollow(followeeUId, followerUId string, action int, ctx context.Context) (err error) {
	// todo  1. 修改follow关系 2. 更新relationship_record表
	//var operation int
	//switch action {
	//case "0":
	//	operation = Invariant
	//case "1":
	//	operation = Increase
	//default:
	//	operation = Invariant
	//}
	now := time.Now().Format("2006-01-02 15:04:05")

	return d.db.Transaction(func(tx *gorm.DB) error {

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

		return nil

	})
}
