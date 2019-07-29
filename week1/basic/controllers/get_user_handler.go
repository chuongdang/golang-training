package controllers

import (
	"basic/dto"
	"basic/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUserGet(ctx *gin.Context) {
	var userGetParam dto.UserGetParam
	if err := ctx.ShouldBindQuery(&userGetParam); err != nil {
		responseError(ctx, err)
		return
	}

	userList, _ := models.ListUser(&userGetParam)

	ctx.JSON(http.StatusOK, &userList)
}