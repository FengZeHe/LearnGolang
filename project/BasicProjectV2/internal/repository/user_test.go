package repository

import (
	"context"
	"database/sql"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/cache"
	cachemocks "github.com/basicprojectv2/internal/repository/cache/mocks"
	"github.com/basicprojectv2/internal/repository/dao"
	daomocks "github.com/basicprojectv2/internal/repository/dao/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestCacheUserRepository_FindByEmail(t *testing.T) {

}

func TestCacheUserRepository_FindById(t *testing.T) {
	testCase := []struct {
		name string
		mock func(ctrl *gomock.Controller) (cache.UserCache, dao.UserDAO)

		ctx context.Context
		id  string

		wantUser domain.User
		wantErr  error
	}{
		{
			name: "查询成功 但未命中缓存",
			mock: func(ctrl *gomock.Controller) (cache.UserCache, dao.UserDAO) {
				id := "123456"
				d := daomocks.NewMockUserDAO(ctrl)
				c := cachemocks.NewMockUserCache(ctrl)
				// 未命中缓存
				c.EXPECT().Get(gomock.Any(), id).
					Return(domain.User{}, cache.ErrKeyNotExist)
				// 查询成功
				d.EXPECT().FindById(gomock.Any(), id).
					Return(dao.User{
						ID:       id,
						Email:    sql.NullString{String: "1@qq.com", Valid: true},
						Password: "123",
						Phone:    sql.NullString{String: "123", Valid: true},
						Birthday: 0,
						Nickname: "",
						Aboutme:  "夏天夏天悄悄过去留下小秘密",
					}, nil)

				// 设置缓存成功
				c.EXPECT().Set(gomock.Any(), domain.User{
					ID:       id,
					Email:    "1@qq.com",
					Password: "123",
					Phone:    "123",
					Aboutme:  "夏天夏天悄悄过去留下小秘密",
				}).Return(nil)
				return c, d
			},
			id:  "123456",
			ctx: context.Background(),
			wantUser: domain.User{
				ID:       "123456",
				Email:    "1@qq.com",
				Password: "123",
				Phone:    "123",
				Aboutme:  "夏天夏天悄悄过去留下小秘密",
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc, ud := tc.mock(ctrl)
			svc := NewCacheUserRepository(ud, uc)
			user, err := svc.FindById(tc.ctx, tc.id)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, user)
		})
	}
}
