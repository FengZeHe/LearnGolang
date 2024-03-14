package logic

import (
	"wiretes/internal/dao"
	"wiretes/internal/model"
)

type UserLogic struct {
	userDao *dao.UserDao
}

func NewUserLogic(userDao *dao.UserDao) *UserLogic {
	return &UserLogic{userDao: userDao}
}

func (logic *UserLogic) CreateUser(user *model.User) error {
	return logic.userDao.Create(user)
}
