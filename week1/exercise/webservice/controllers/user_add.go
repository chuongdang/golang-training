package controllers

import (
	"net/http"
	"webservice/dto"
	"webservice/models/users"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (uc *UserController) HandleAddUserRequest(ctx *gin.Context) {
	var user *dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	users.AddUser(user)

	responseSuccess(ctx, user)
}
