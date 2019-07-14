package models

import (
	"fmt"
	"github.com/chuongdang/golang-libs/db"
	"product-service/constants"
	"product-service/dto"
	"product-service/entities"
)

func GetListProduct(params *dto.ProductGetParams) (products []entities.Product, err error) {
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = constants.PRODUCT_DEFAULT_LIMIT
	}

	offset := (params.Page - 1) * params.Limit

	err = db.GetConn().Select(
		&products,
		fmt.Sprintf(
			"SELECT id_product, fk_category, name, image FROM product LIMIT %v, %v",
			offset,
			params.Limit,
		),
	)
	return
}
