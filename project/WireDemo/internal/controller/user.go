package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wiretes/internal/logic"
	"wiretes/internal/model"
)

type UserController struct {
	userLogic *logic.UserLogic
}

func NewUserController(userLogic *logic.UserLogic) *UserController {
	return &UserController{userLogic: userLogic}
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := controller.userLogic.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "create success",
	})
}
