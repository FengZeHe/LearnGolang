package dao

import (
	"context"
	"gorm.io/gorm"
	"prometheusdemo/internal/domain"
)

type UserDAO interface {
	GetUser(ctx context.Context) (domain.User, error)
}

type GORMUserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}

// ToDo 改 GormUserDAO
func (u *GORMUserDAO) GetUser(ctx context.Context) (domain.User, error) {
	return domain.User{}, nil
}
