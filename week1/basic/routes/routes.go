package routes

import (
	"basic/controllers"
	"basic/middlewares"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.Auth)

	router.GET("/user", controllers.HandleUserGet)
	router.POST("/user", controllers.HandleUserCreate)
}

func Start() {
	router.Run(":6969")
}
