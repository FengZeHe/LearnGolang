package service

import (
	"context"
	"database/sql"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
	repomocks "github.com/basicprojectv2/internal/repository/mocks"
	"github.com/basicprojectv2/pkg/bcrypt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
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
		email    string
		password string

		wantUser domain.User
		wantErr  error
	}{
		{
			name: "登录成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)

				repo.EXPECT().
					FindByEmail(gomock.Any(), sql.NullString{String: "1@qq.com", Valid: true}.String).
					Return(
						domain.User{
							Email:    sql.NullString{String: "1@qq.com", Valid: true},
							Password: "$2a$10$GFSbeRa5RVX912wj0SGZHuGjeJuDEB3eOlXDv8GZ/yAaP4rKY9Roq",
							Phone:    sql.NullString{String: "12345", Valid: true},
						}, nil)
				return repo
			},
			email:    "1@qq.com",
			password: "123",

			wantUser: domain.User{
				Email:    sql.NullString{String: "1@qq.com", Valid: true},
				Password: "$2a$10$GFSbeRa5RVX912wj0SGZHuGjeJuDEB3eOlXDv8GZ/yAaP4rKY9Roq",
				Phone:    sql.NullString{String: "12345", Valid: true},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := tc.mock(ctrl)
			svc := NewUserService(repo)
			user, err := svc.Login(tc.ctx, tc.email, tc.password)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, user)
		})
	}

}
