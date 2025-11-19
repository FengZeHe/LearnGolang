package dao

import (
	"context"

	"gorm.io/gorm"
)

type GORMInteractive struct {
	db *gorm.DB
}

type InteractiveDAO interface {
	AddReadCount(ctx context.Context, id string) (err error)
}

func NewInteractiveDAO(db *gorm.DB) InteractiveDAO {
	return &GORMInteractive{db: db}
}

func (i *GORMInteractive) AddReadCount(ctx context.Context, id string) (err error) {
	//i.db.Model(domain.Interactive{}).Table("interactive")

	return nil
}
