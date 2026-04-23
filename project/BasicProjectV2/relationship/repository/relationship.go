package repository

import "github.com/basicprojectv2/relationship/repository/dao"

type relationshipRepository struct {
	dao dao.RelationshipDAO
}

type RelationshipRepository interface {
}

func NewRelationshipRepository(dao dao.RelationshipDAO) RelationshipRepository {
	return &relationshipRepository{dao: dao}
}
