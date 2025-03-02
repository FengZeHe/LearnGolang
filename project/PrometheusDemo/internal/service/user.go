package service

import "prometheusdemo/internal/repository"

type UserService interface {
	GetUser()
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (u userService) GetUser() {
	//TODO implement me
	panic("implement me")
}
