package controllers

import (
	"webservice/dto"
	"webservice/models/users"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) HandleUserListRequest(ctx *gin.Context) {
	users, err := users.GetListUser()

	if err != nil {
		ctx.JSON(500, dto.ErrorResponse{Message: err.Error()})
		return
	}

	uc.Response(ctx, users)
}
