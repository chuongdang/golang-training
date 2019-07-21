package controllers

import (
	"basic/version"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUserNameRequest(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.String(http.StatusOK, "Hello %s. This is version %s", name, version.GetVersion())
}