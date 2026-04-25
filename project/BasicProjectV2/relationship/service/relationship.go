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
