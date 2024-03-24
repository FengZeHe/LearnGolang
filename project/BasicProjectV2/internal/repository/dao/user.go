package dao

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"gorm.io/gorm"
)

type GORMUserDAO struct {
	db *gorm.DB
}

type UserDAO interface {
	Insert(ctx context.Context, u domain.User) error
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}

func (dao *GORMUserDAO) Insert(ctx context.Context, u domain.User) error {
	return nil
}
