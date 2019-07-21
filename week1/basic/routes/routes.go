package routes

import (
	"basic/controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.Default()

	router.POST("/user", controllers.HandleUserCreate)
}

func Start() {
	router.Run(":6969")
}
