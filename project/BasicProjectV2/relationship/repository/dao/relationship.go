package dao

import "gorm.io/gorm"

type GORMRelationship struct {
	db *gorm.DB
}

type RelationshipDAO interface {
}

func NewGORMRelationshipDAO(db *gorm.DB) RelationshipDAO {
	return &GORMRelationship{db: db}
}
