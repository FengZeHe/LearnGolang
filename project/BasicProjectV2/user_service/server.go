package main

import (
	"context"
	"github.com/basicprojectv2/user_service/domain"
	"github.com/basicprojectv2/user_service/pkg/jwt"
	"github.com/basicprojectv2/user_service/service"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *UserService) UserRegister(ctx context.Context, req *UserRegisterReq) (signResp *UserRegisterResp, err error) {
	signResp = &UserRegisterResp{} // 初始化signResp

	err = s.svc.Signup(ctx, domain.User{Email: req.Email, Password: req.Password})
	if err != nil {
		signResp.Msg = err.Error()
		return signResp, err
	}
	signResp.Msg = "注册成功"

	return signResp, nil
}

func (s *UserService) UserLogin(ctx context.Context, req *UserLoginReq) (loginResp *UserLoginResp, err error) {
	loginResp = &UserLoginResp{}
	u, err := s.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		log.Println(err)
		loginResp.Msg = err.Error()
		return loginResp, err
	}

	token, err := jwt.GenToken(u.ID)
	if err != nil {
		log.Println(err)
		loginResp.Msg = err.Error()
		return nil, err
	}
	loginResp.Token = token
	return loginResp, nil

}

func (s *UserService) Hi(ctx context.Context, in *emptypb.Empty) (resp *HiResp, err error) {
	resp = &HiResp{}
	resp.Msg = "Hi! 这里是user_service"
	return resp, nil
}
