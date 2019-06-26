package transformer

import (
	"product-service/dto"
	"product-service/entities"
)

func MapProductForResponse(
	productList *[]entities.Product,
	categoryMap map[int]dto.Category,
) (result []dto.Product, err error) {
	for _, product := range *productList {
		productDto := dto.Product{
			ID:       product.ID,
			Name:     product.Name,
			Image:    product.Image,
			Category: categoryMap[product.FkCategory],
		}
		result = append(result, productDto)
	}
	return result, nil
}
