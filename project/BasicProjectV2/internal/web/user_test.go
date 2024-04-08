package web

import (
	"bytes"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	svcmocks "github.com/basicprojectv2/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_Login(t *testing.T) {
	testCases := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) (service.UserService, service.CodeService)
		email    string
		password string

		reqBuilder func(t *testing.T) *http.Request

		// 预期输出
		wantCode int
		wantBody string
	}{
		{
			name: "登录成功",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				userSvc := svcmocks.NewMockUserService(ctrl)
				userSvc.EXPECT().
					Login(gomock.Any(), "1@qq.com", "123").
					Return(domain.User{}, nil)
				codeSvc := svcmocks.NewMockCodeService(ctrl)
				return userSvc, codeSvc
			},
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/v2/users/login", bytes.NewReader(
					[]byte(`{"email":"1@qq.com","password":"123"}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			wantCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// 登录函数
			userSvc, codeSvc := tc.mock(ctrl)
			hdl := NewUserHandler(userSvc, codeSvc)

			// 注册路由
			serve := gin.Default()
			hdl.RegisterRoutes(serve)

			// 准备Req和记录的recorder
			req := tc.reqBuilder(t)
			recorder := httptest.NewRecorder()

			//执行
			serve.ServeHTTP(recorder, req)

			//断言
			assert.Equal(t, tc.wantCode, recorder.Code)
		})
	}

}
