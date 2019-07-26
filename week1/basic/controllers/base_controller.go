package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func responseError(ctx *gin.Context, err error) {
	ctx.String(http.StatusInternalServerError, "Error %s", err.Error())
}