package repository

import (
	"context"
	"encoding/json"
	"log"
	"prometheusdemo/internal/domain"
	"prometheusdemo/internal/repository/cache"
	"prometheusdemo/internal/repository/dao"
)

type UserRepository interface {
	GetUser() (users []domain.User, err error)
	GetUserCache() (users []domain.User, err error)
}

type userRepository struct {
	dao dao.UserDAO
	c   cache.UserCache
}

func NewCacheUserRepository(dao dao.UserDAO, c cache.UserCache) UserRepository {
	return &userRepository{
		dao: dao,
		c:   c,
	}
}

func (u userRepository) GetUser() (users []domain.User, err error) {
	users, err = u.dao.GetUser()
	if err != nil {
		return users, err
	}

	jsonStr, err := json.Marshal(users)
	if err = u.c.Set(context.Background(), "users", string(jsonStr)); err != nil {
		log.Println(err)
		log.Println("set cache error", err)
	}

	return users, nil
}

func (u userRepository) GetUserCache() (users []domain.User, err error) {
	data, err := u.c.Get(context.Background(), "users")
	if err != nil {
		return users, err
	}
	return data, err
}
