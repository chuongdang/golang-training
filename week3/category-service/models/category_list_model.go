package models

import (
	"category-service/dberr"
	"category-service/entities"
	"fmt"
	"strings"

	"github.com/chuongdang/golang-libs/db"
)

func GetListCategory(idList []string) (categories []entities.Category, err error) {
	idListString := strings.Join(idList, ",")
	query := "SELECT id_category, name, image FROM category"
	if len(idList) > 0 {
		query += fmt.Sprintf(" WHERE id_category IN (%v)", idListString)
	}
	err = db.GetConn().Select(&categories, query)
	if err != nil {
		dberr.ErrChannel <- true
		err = db.GetConn().Select(&categories, query)
	}
	return
}
