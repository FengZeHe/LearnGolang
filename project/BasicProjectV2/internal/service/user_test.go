package service

import (
	"context"
	"database/sql"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
	repomocks "github.com/basicprojectv2/internal/repository/mocks"
	"github.com/basicprojectv2/pkg/bcrypt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

// 测试password加密功能
func TestPasswordEncrypt(t *testing.T) {
	password := []byte("123")
	encrypted, err := bcrypt.GetPwd(string(password))
	assert.NoError(t, err)
	result := bcrypt.ComparePwd(encrypted, string([]byte("123")))
	if result != true {
		err = errors.New("验证错误")
	}
	assert.NoError(t, err)
}

// 测试登录功能
func Test_userService_Login(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) repository.UserRepository

		// 预期输入
		ctx      context.Context
		email    sql.NullString
		password string

		//实际值
		wantUser domain.User
		wantErr  error
	}{
		{
			name: "登录成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				// 使用repomocks包创建一个模拟的UserRepository
				repo := repomocks.NewMockUserRepository(ctrl)
				//	设置模拟UserRepository的FindByEmail方法的预期行为，期望FindByEmail方法被调用时，传入的邮箱为 "1@qq.com"
				// 返回值是一个具有特定字段值的 domain.User对象
				repo.EXPECT().
					FindByEmail(gomock.Any(), "1@qq.com").
					Return(domain.User{
						Email:    "1@qq.com",
						Password: "$2a$10$GFSbeRa5RVX912wj0SGZHuGjeJuDEB3eOlXDv8GZ/yAaP4rKY9Roq",
						Phone:    "123",
					}, nil)
				return repo
			},
			email:    sql.NullString{String: "1@qq.com", Valid: true},
			password: "123",

			wantUser: domain.User{
				Email:    "1@qq.com",
				Password: "$2a$10$GFSbeRa5RVX912wj0SGZHuGjeJuDEB3eOlXDv8GZ/yAaP4rKY9Roq",
				Phone:    "123",
			},
		},
		{
			name: "密码或用户名错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().
					FindByEmail(gomock.Any(), "1@qq.com").
					Return(
						domain.User{
							Email:    "1@qq.com",
							Password: "$2a$10$GFSbeRa5RVX912wj0SGZHuGjeJuDEB3eOlXDv8GZ/yAaP4rKY9Roq",
							Phone:    "123",
						}, nil)
				return repo
			},
			email:    sql.NullString{String: "1@qq.com", Valid: true},
			password: "123123123",
			wantErr:  ErrInvalidUserOrPassword,
		},
		{
			name: "用户不存在",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().
					FindByEmail(gomock.Any(), "xx@qq.com").
					Return(domain.User{}, repository.ErrUserNotFound)
				return repo
			},
			email:    sql.NullString{String: "xx@qq.com", Valid: true},
			password: "xxxxxxxx",
			//wantUser: domain.User{},
			wantErr: ErrInvalidUserOrPassword,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// 模拟UserRepository 来创建 UserSerivce
			repo := tc.mock(ctrl)
			svc := NewUserService(repo)

			// 调用Login函数
			user, err := svc.Login(tc.ctx, tc.email.String, tc.password)
			assert.Equal(t, tc.wantUser, user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

// 测试注册功能
func Test_userService_Signup(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) repository.UserRepository

		// 预期输入
		ctx  context.Context
		user domain.User
		//实际值
		wantErr error
	}{
		{
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil)
				return repo
			},
			wantErr: nil,
		},
		{
			name: "加密失败",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(bcrypt.ErrGenPasswd)
				return repo
			},
			wantErr: bcrypt.ErrGenPasswd,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := tc.mock(ctrl)
			svc := NewUserService(repo)
			err := svc.Signup(tc.ctx, tc.user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
