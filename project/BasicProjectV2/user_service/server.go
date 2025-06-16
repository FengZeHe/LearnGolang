package user_service

import (
	"context"
	"github.com/basicprojectv2/internal/service"
	"log"
)

type UserService struct {
	UnimplementedUserServiceServer
	svc service.UserService
}

// 新建构造函数
func NewUerService(svc service.UserService) *UserService {
	return &UserService{svc: svc}
}

func (s *UserService) GetUserById(ctx context.Context, req *GetUserByIdReq) (*User, error) {
	user, err := s.svc.FindById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	log.Printf("user: %v", user)
	return &User{
		Id:       user.ID,
		Email:    user.Email,
		Phone:    user.Phone,
		Birthday: int32(user.Birthday),
		Nickname: user.Nickname,
		Aboutme:  user.Aboutme,
		Role:     user.Role,
	}, nil
}
