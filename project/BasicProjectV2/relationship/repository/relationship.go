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
	HandleBlock(uid string, req domain.BlockReq, ctx context.Context) error
	QueryRelationship(uid, targetUid string, ctx context.Context) (userStatus domain.UserStatus, err error)
	QueryFolloweeList(uid string, req domain.FollowListReq, ctx context.Context) (resp domain.FolloweeListResp, err error)
	QueryFollowerList(uid string, req domain.FollowListReq, ctx context.Context) (resp domain.FollowerListResp, err error)
	CountRelationship(uid string, ctx context.Context) (resp domain.RelationshipCount, err error)
}

func NewRelationshipRepository(dao dao.RelationshipDAO) RelationshipRepository {
	return &relationshipRepository{dao: dao}
}
func (r *relationshipRepository) HandleFollow(uid string, req domain.FollowReq, ctx context.Context) (err error) {
	if err = r.dao.HandleFollow(uid, req.Uid, req.Action, ctx); err != nil {
		return err
	}
	return nil
}

func (r *relationshipRepository) HandleBlock(uid string, req domain.BlockReq, ctx context.Context) (err error) {
	if err = r.dao.HandleBlock(uid, req.Uid, req.Action, ctx); err != nil {
		return err
	}
	return nil
}

func (r *relationshipRepository) QueryRelationship(uid, targetUid string, ctx context.Context) (userStatus domain.UserStatus, err error) {
	userStatus, err = r.dao.QueryRelationship(uid, targetUid, ctx)
	if err != nil {
		return userStatus, err
	}
	return userStatus, nil
}

func (r *relationshipRepository) QueryFolloweeList(uid string, req domain.FollowListReq, ctx context.Context) (userStatus domain.FolloweeListResp, err error) {
	return r.dao.QueryFolloweeList(uid, req.PageIndex, req.PageSize, ctx)
}

func (r *relationshipRepository) QueryFollowerList(uid string, req domain.FollowListReq, ctx context.Context) (resp domain.FollowerListResp, err error) {
	return r.dao.QueryFollowerList(uid, req.PageIndex, req.PageSize, ctx)
}

func (r *relationshipRepository) CountRelationship(uid string, ctx context.Context) (resp domain.RelationshipCount, err error) {
	return r.dao.CountRelationship(uid, ctx)
}
