package users

import (
	"webservice/dto"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetListUser() ([]dto.User, error) {
	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/nordic")
	if err != nil {
		return nil, err
	}

	var users []dto.User
	db.Select(&users, "SELECT * FROM users")

	return users, nil
}
