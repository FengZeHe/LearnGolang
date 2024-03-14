package dao

import (
	"gorm.io/gorm"
	"wiretes/internal/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// Create 方法
func (dao *UserDao) Create(user *model.User) error {
	return dao.db.Create(user).Error
}
