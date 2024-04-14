package dao

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

// 测试dao层 Insert函数
func TestGORMUserDAO_Insert(t *testing.T) {
	testCases := []struct {
		name string
		mock func(t *testing.T) *sql.DB
		user User
		ctx  context.Context

		wantErr error
	}{
		{
			name: "写入成功",
			mock: func(t *testing.T) *sql.DB {
				db, mock, err := sqlmock.New()
				assert.NoError(t, err)
				mockRes := sqlmock.NewResult(123, 1)
				// 这边要求传入的是 sql 的正则表达式
				mock.ExpectExec("INSERT INTO .*").
					WillReturnResult(mockRes)
				return db
			},
			ctx:  context.Background(),
			user: User{Aboutme: "夏天夏天悄悄过去"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			sqlDB := tc.mock(t)
			db, err := gorm.Open(mysql.New(mysql.Config{
				Conn:                      sqlDB,
				SkipInitializeWithVersion: true,
			}), &gorm.Config{
				DisableAutomaticPing:   true,
				SkipDefaultTransaction: true,
			})
			assert.NoError(t, err)
			dao := NewUserDAO(db)
			err = dao.Insert(tc.ctx, tc.user)
			assert.Equal(t, tc.wantErr, err)

		})
	}
}
