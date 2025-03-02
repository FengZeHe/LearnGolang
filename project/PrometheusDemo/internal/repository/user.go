package repository

import "prometheusdemo/internal/repository/dao"

type UserRepository interface {
	GetUser()
}

type userRepository struct {
	dao dao.UserDAO
}

func NewCacheUserRepository(dao dao.UserDAO) UserRepository {
	return &userRepository{
		dao: dao,
	}
}

func (u userRepository) GetUser() {
	//TODO implement me
	panic("implement me")
}
