package controller

import (
	"BasicProject/logic"
	"BasicProject/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// SignIn 模拟logic.SignIn函数
func (m *MockSignInLogic) Sign(user *models.RegisterFormByEmail) error {
	return nil
}

type MockSignInLogic struct{}

// SignIn 模拟logic.SignIn函数
func (m *MockSignInLogic) SignIn(user *models.RegisterFormByEmail) error {
	// 在这里模拟逻辑，例如直接返回nil表示注册成功
	return nil
}

func TestHandleUserSignIn_Success(t *testing.T) {
	// 创建一个模拟的HTTP请求
	form := models.RegisterFormByEmail{Email: "1@qq.com", Password: "123", ConfirmPassword: "123"}
	reqBody, err := json.Marshal(form)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/api/v1/signin", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建一个response recorder来记录响应
	rec := httptest.NewRecorder()

	// 创建一个gin context
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	// 设置logic层的mock
	logic.SetSignInLogic(&logic.MockSignInLogic{})

	// 调用controller层处理函数
	HandleUserSignIn(ctx)

	// 检查响应状态码
	assert.Equal(t, http.StatusOK, rec.Code)

	// 检查响应内容
	//var response map[string]interface{}
	//err = json.Unmarshal(rec.Body.Bytes(), &response)
	//assert.NoError(t, err)
	//assert.Equal(t, "success", response["message"])
}
