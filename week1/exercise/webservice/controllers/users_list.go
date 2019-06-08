package controllers

import (
	"webservice/dto"
	"webservice/models/users"

	"github.com/gin-gonic/gin"
)

func HandleUserListRequest(ctx *gin.Context) {
	users, err := users.GetListUser()

	if err != nil {
		ctx.AbortWithStatusJSON(500, dto.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(200, dto.Response{Data: users})
}
