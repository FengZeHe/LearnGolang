package service

import (
	"github.com/basicprojectv2/relationship/repository"
)

type relationshipService struct {
	r repository.RelationshipRepository
}

type RelationshipService interface {
}

func NewRelationshipService(r repository.RelationshipRepository) RelationshipService {
	return &relationshipService{r: r}
}
