package controllers

import (
	"basic/db"
	"basic/dto"
	"basic/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUserCreate(ctx *gin.Context) {
	var userAddParam dto.UserAddParam
	if err := ctx.ShouldBindJSON(&userAddParam); err != nil {
		responseError(ctx, err)
		return
	}

	transaction, _ := db.GetConn().Beginx()
	user, _ := models.GetUser(transaction, &userAddParam)
	if user != nil {
		responseError(ctx, errors.New(fmt.Sprintf("User %s is duplicated", userAddParam.Email)))
		return
	}

	if err := models.AddUser(transaction, &userAddParam); err != nil {
		responseError(ctx, err)
		transaction.Rollback()
		return
	}

	transaction.Commit()

	ctx.JSON(http.StatusOK, &userAddParam)
}