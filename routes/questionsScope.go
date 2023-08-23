package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddQuestionRoutes(rg *gin.RouterGroup) {
	questions := rg.Group("/")

	questions.GET("/questions", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "questions")
	})
}
