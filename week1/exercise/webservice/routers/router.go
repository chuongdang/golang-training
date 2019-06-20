package routers

import (
	"webservice/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	userController := controllers.UserController{}

	router.GET("/ping", controllers.HandlePingRequest)
	router.GET("/users", userController.HandleUserListRequest)
	router.GET("/users/:id", userController.HandleUserDetailRequest)
	router.POST("/users", userController.HandleAddUserRequest)
}

/*
Run run application
*/
func Run(port string) {
	router.Run(":" + port)
}
