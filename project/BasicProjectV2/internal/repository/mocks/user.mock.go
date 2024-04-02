package repomocks

import (
	context2 "context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"reflect"
)

type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockrecorder
}

func (m *MockUserRepository) Create(ctx context.Context, u domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

func (m *MockUserRepository) FindByPhone(ctx context2.Context, phone string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPhone", ctx, phone)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

type MockUserRepositoryMockrecorder struct {
	mock *MockUserRepository
}

func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockrecorder{mock}
	return mock
}

func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockrecorder {
	return m.recorder
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockrecorder) FindByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), ctx, email)
}
