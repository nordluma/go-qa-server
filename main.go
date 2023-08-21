package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func addQuestionRoutes(rg *gin.RouterGroup) {
	questions := rg.Group("/questions")

	questions.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "questions")
	})
}

func setupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	routes := r.Group("/v1")
	routes.GET("/health_check", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	addQuestionRoutes(routes)

	return r
}

func main() {
	r := setupRoutes()

	r.Run(":8080")
}
