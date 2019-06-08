package routers

import (
	"webservice/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.GET("/ping", controllers.HandlePingRequest)
	router.GET("/users", controllers.HandleUserListRequest)
}

/*
Run run application
*/
func Run(port string) {
	router.Run(":" + port)
}
