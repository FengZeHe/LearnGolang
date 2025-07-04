package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/change", func(context *gin.Context) {})

}
