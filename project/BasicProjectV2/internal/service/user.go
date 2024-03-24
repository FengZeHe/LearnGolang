package service

import "github.com/basicprojectv2/internal/repository"

type UserService interface {
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
