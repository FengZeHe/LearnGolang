package service

import (
	"context"

	"github.com/basicprojectv2/relationship/domain"
	"github.com/basicprojectv2/relationship/repository"
)

type relationshipService struct {
	repo repository.RelationshipRepository
}

type RelationshipService interface {
	HandleFollow(uid string, req domain.FollowReq, ctx context.Context) error
	HandleBlock(uid string, req domain.BlockReq, ctx context.Context) error
	QueryRelationship(uid, targetUid string, ctx context.Context) (userStatus domain.UserStatus, err error)
	QueryFolloweeList(uid string, req domain.FollowListReq, ctx context.Context) (resp domain.FolloweeListResp, err error)
	QueryFollowerList(uid string, req domain.FollowListReq, ctx context.Context) (resp domain.FollowerListResp, err error)
	CountRelationship(uid string, ctx context.Context) (resp domain.RelationshipCount, err error)
}

func NewRelationshipService(r repository.RelationshipRepository) RelationshipService {
	return &relationshipService{repo: r}
}

func (r *relationshipService) HandleFollow(uid string, req domain.FollowReq, ctx context.Context) (err error) {
	if err = r.repo.HandleFollow(uid, req, ctx); err != nil {
		return err
	}
	return nil
}

func (r *relationshipService) HandleBlock(uid string, req domain.BlockReq, ctx context.Context) (err error) {
	if err = r.repo.HandleBlock(uid, req, ctx); err != nil {
		return err
	}
	return nil
}

func (r *relationshipService) QueryRelationship(uid, targetUid string, ctx context.Context) (userStatus domain.UserStatus, err error) {
	return r.repo.QueryRelationship(uid, targetUid, ctx)
}

func (r *relationshipService) QueryFolloweeList(uid string, req domain.FollowListReq, ctx context.Context) (resp domain.FolloweeListResp, err error) {
	return r.repo.QueryFolloweeList(uid, req, ctx)
}

func (r *relationshipService) QueryFollowerList(uid string, req domain.FollowListReq, ctx context.Context) (resp domain.FollowerListResp, err error) {
	return r.repo.QueryFollowerList(uid, req, ctx)
}

func (r *relationshipService) CountRelationship(uid string, ctx context.Context) (resp domain.RelationshipCount, err error) {
	return r.repo.CountRelationship(uid, ctx)
}
