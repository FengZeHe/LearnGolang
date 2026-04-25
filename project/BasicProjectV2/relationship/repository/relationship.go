package repository

import (
	"context"

	"github.com/basicprojectv2/relationship/domain"
	"github.com/basicprojectv2/relationship/repository/dao"
)

type relationshipRepository struct {
	dao dao.RelationshipDAO
}

type RelationshipRepository interface {
	HandleFollow(uid string, req domain.FollowReq, ctx context.Context) error
}

func NewRelationshipRepository(dao dao.RelationshipDAO) RelationshipRepository {
	return &relationshipRepository{dao: dao}
}
func (r *relationshipRepository) HandleFollow(uid string, req domain.FollowReq, ctx context.Context) (err error) {
	if err = r.dao.HandleFollow(uid, req.Uid, req.Follow, ctx); err != nil {
		return err
	}
	return nil
}
