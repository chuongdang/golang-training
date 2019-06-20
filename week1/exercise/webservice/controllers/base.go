package controllers

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (bc *BaseController) Response(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, data)
}

func responseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, data)
}
