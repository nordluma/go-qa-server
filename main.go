package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nordluma/go-qa-server/routes"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/health_check", healthCheck)

	routesV1 := r.Group("/v1")
	routes.AddQuestionRoutes(routesV1)

	return r
}

func healthCheck(ctx *gin.Context) {
	ctx.Status(200)
}

func main() {
	r := setupRoutes()

	r.Run(":8080")
}
