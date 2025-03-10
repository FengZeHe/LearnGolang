package service

import (
	"prometheusdemo/internal/domain"
	"prometheusdemo/internal/repository"
)

type UserService interface {
	GetUser() (users []domain.User, err error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (u userService) GetUser() (users []domain.User, err error) {
	users, err = u.repo.GetUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}
