package controllers

import (
	"category-service/models"
	"github.com/gin-gonic/gin"
	"strings"
)

func HandleCategoryList(ctx *gin.Context) {
	idList := ctx.Query("ids")
	idListArray := strings.Split(idList, ",")
	products, err := models.GetListCategory(idListArray)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not get Category List. Err:" + err.Error(),
		})
		return
	}
	ctx.JSON(200, products)
}
