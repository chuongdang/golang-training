package models

import (
	"basic/db"
	"basic/dto"
	"basic/entity"
)

func ListUser(param *dto.UserGetParam) (userList []entity.User, err error) {
	offset := (param.Page - 1) * param.Limit

	query := "SELECT id_user, fk_role, name, email FROM user LIMIT ?,?"

	rows, err := db.GetConn().Queryx(query, offset, param.Limit)
	if err != nil {
		return
	}

	for rows.Next() {
		var user entity.User
		err = rows.StructScan(&user)
		if err != nil {
			return
		}
		userList = append(userList, user)
	}

	return
}