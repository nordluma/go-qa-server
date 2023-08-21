package main

import (
	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/health_check", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	return r
}

func main() {
	r := setupRoutes()

	r.Run(":8080")
}
