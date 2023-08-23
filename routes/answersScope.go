package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAnswerRoutes(rg *gin.RouterGroup) {
	questions := rg.Group("/")

	questions.GET("/answers", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "answers")
	})
}
