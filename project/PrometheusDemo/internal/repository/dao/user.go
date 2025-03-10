package dao

import (
	"gorm.io/gorm"
	"prometheusdemo/internal/domain"
)

type UserDAO interface {
	GetUser() ([]domain.User, error)
}

type GORMUserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}

func (u *GORMUserDAO) GetUser() ([]domain.User, error) {
	var data []domain.User
	if err := u.db.Table("users").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
