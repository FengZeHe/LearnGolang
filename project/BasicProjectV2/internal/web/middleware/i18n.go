package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func I18nMiddleware(i18nBundle *i18n.Bundle) gin.HandlerFunc {
	return func(c *gin.Context) {
		acceptLang := c.GetHeader("Accept-Language")
		if acceptLang == "" {
			acceptLang = "zh-CN" //默认语言
		}

		localizer := i18n.NewLocalizer(i18nBundle, acceptLang)
		c.Set("localizer", localizer)
		c.Next()
	}
}
