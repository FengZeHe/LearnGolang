package service

import (
	"context"
	cb "limiterdemo/user_service/circuitbraker"
	service "limiterdemo/user_service/proto/user_service"
)

type UserService struct {
	service.UnimplementedUserServiceServer
	circuitBreaker *cb.CircuitBreaker
}

func NewUserService(cb *cb.CircuitBreaker) *UserService {
	return &UserService{
		circuitBreaker: cb,
	}
}

func (s *UserService) GetUserById(ctx context.Context, req *service.GetUserReq) (*service.User, error) {
	userId := req.Id
	user := &service.User{
		Userid:   userId,
		Username: "熊二",
		Age:      "8",
	}
	return user, nil
}

// 设置熔断器工作状态
func (s *UserService) ControlCircuitBraker(ctx context.Context, req *service.CircuitBrakerReq) (*service.CircuitBrakerResp, error) {
	mReq := req.Manual
	eReq := req.Enabled
	switch mReq {
	case 0:
		s.circuitBreaker.DisableManul()
	case 1:
		s.circuitBreaker.ManualClose()
	case 2:
		s.circuitBreaker.ManualOpen()
	}

	switch eReq {
	case 0:
		s.circuitBreaker.Enable()
	case 1:
		s.circuitBreaker.Disable()
	}

	//log.Println(req)

	return nil, nil
}

func (s *UserService) CoreBusiness(ctx context.Context) (*service.CoreResp, error) {
	return &service.CoreResp{
		Msg: "都jb哥们",
	}, nil
}
