package routes

import (
	"category-service/config"
	"category-service/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	port := os.Getenv("CATEGORY_SERVICE_PORT")
	if port == "" {
		port = config.Config.Api.Port
	}
	route := gin.Default()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.GET("/categories", controllers.HandleCategoryList)
	route.Run(":" + port)
}
