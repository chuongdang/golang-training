package models

import (
	"product-service/entities"

	"github.com/chuongdang/golang-libs/db"
)

func GetListProduct() (products []entities.Product, err error) {
	err = db.GetConn().Select(&products, "SELECT id_product, fk_category, name, image FROM product")
	return
}
