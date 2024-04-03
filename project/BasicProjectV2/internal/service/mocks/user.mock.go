package svcmocks

import "github.com/golang/mock/gomock"

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}
