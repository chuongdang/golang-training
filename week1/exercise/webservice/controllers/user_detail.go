package controllers

import (
	"webservice/dto"
	"webservice/models/users"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) HandleUserDetailRequest(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := users.GetUserById(userId)

	if err != nil {
		errMsg := dto.ErrorResponseWithUserID{
			UserId: userId,
		}
		errMsg.Message = err.Error()
		errMsg.Data = "Test"
		errMsg.ShowData()
		ctx.JSON(500, errMsg)
		return
	}

	uc.Response(ctx, user)
}
