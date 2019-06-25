package models

import (
	"product-service/db"
	"product-service/entities"
)

func GetListProduct() (products []entities.Product, err error) {
	err = db.GetConn().Select(&products, "SELECT id_product, fk_category, name, image FROM product")
	return
}
