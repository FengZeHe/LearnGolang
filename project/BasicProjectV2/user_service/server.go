package user_service

import "github.com/basicprojectv2/internal/repository"

type UserService struct {
	UnimplementedUserServiceServer
	repo repository.UserRepository
}

// 新建构造函数
func NewUerService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
