package models

import (
	"encoding/json"
	"log"
	"product-service/dto"
	"product-service/entities"
	"product-service/service/category_service"
	"strconv"
)

func GetCategoriesDetailFromProductList(
	productList *[]entities.Product,
) (categoryMap map[int]dto.Category, err error) {
	categoryMap = make(map[int]dto.Category)
	var categoryIdList []string
	for _, product := range *productList {
		categoryIdList = append(categoryIdList, strconv.Itoa(product.FkCategory))
	}

	var categoryList []dto.Category
	categoryData, _ := category_service.FetchCategoryByIdList(categoryIdList)
	err = json.Unmarshal(categoryData, &categoryList)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, category := range categoryList {
		categoryMap[int(category.ID)] = category
	}
	return
}
