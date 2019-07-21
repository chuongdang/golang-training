package models

import (
	"basic/db"
	"basic/dto"
	"fmt"
)

func AddUser(param *dto.UserAddParam) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (`%s`, `%s`, `%s`, `%s`)" +
			"VALUES (%d, \"%s\", \"%s\", \"%s\")",
		"fk_role",
		"name",
		"email",
		"password",
		param.FkRole,
		param.Name,
		param.Email,
		param.Password,
	)
	if _, err := db.GetConn().Exec(query); err != nil {
		return err
	}
	return nil
}