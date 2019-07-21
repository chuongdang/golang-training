package controllers

import (
	"basic/dto"
	"basic/models"
	"github.com/gin-gonic/gin"
)

func HandleUserCreate(ctx *gin.Context) {
	var userAddParam dto.UserAddParam
	if err := ctx.ShouldBindJSON(&userAddParam); err != nil {
		ctx.String(500, "Shit happens! Error %s", err.Error())
	}

	if err := models.AddUser(&userAddParam); err != nil {
		ctx.String(500, "Shit happens! Error %s", err.Error())
	}
}