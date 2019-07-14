package controllers

import (
	"log"
	"product-service/dto"
	"product-service/models"
	"product-service/transformer"

	"github.com/gin-gonic/gin"
)

func HandleProductList(ctx *gin.Context) {
	params := dto.ProductGetParams{}
	if err := ctx.BindQuery(&params); err != nil {
			log.Println(err)
	}
	products, err := models.GetListProduct(&params)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not get Product List. Err:" + err.Error(),
		})
		return
	}

	categoriesData, err := models.GetCategoriesDetailFromProductList(&products)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not get Category Info. Err:" + err.Error(),
		})
		return
	}
	productsDto, _ := transformer.MapProductForResponse(&products, categoriesData)

	ctx.JSON(200, productsDto)
}
