package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nordluma/go-qa-server/routes"

	_ "github.com/lib/pq"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/health_check", healthCheck)

	routesV1 := r.Group("/v1")
	routes.AddQuestionRoutes(routesV1)
	routes.AddAnswerRoutes(routesV1)

	return r
}

func healthCheck(ctx *gin.Context) {
	ctx.Status(200)
}

func main() {
	r := setupRoutes()

	_, err := sqlx.Connect(
		"postgres",
		"postgres://postgres:password@localhost/qadb",
	)
	if err != nil {
		log.Fatal("Could not establis connection to database: ", err)
	}

	r.Run(":8080")
}
