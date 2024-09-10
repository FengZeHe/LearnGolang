package web

import (
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
)

type DraftHandler struct {
	svc service.DraftService
}

func NewDraftHandler(svc service.DraftService) *DraftHandler {
	return &DraftHandler{svc: svc}
}

func (r *DraftHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	rg := server.Group("/v2/draft/")
	rg.Use(loginCheck)
	rg.GET("/getArticles", r.GetArticles)
	rg.POST("/addArticle", r.AddArticle)
	rg.POST("/updateArticle", r.AddArticle)
}

func (r *DraftHandler) GetArticles(c *gin.Context) {

}

func (r *DraftHandler) AddArticle(c *gin.Context) {
	// todo 检查请求参数是否正确

	// todo 请求参数正确，往service层执行
}

func (r *DraftHandler) UpdateArticle(c *gin.Context) {

}
