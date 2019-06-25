package routes

import (
	"os"
	"product-service/config"
	"product-service/controllers"

	"github.com/gin-gonic/gin"
)

func Start() {
	port := os.Getenv("PRODUCT_SERVICE_PORT")
	if port == "" {
		port = config.PORT
	}
	route := gin.Default()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.GET("/products", controllers.HandleProductList)
	route.Run(":" + port)
}
