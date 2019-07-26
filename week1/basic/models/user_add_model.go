package models

import (
	"basic/dto"
	"basic/entity"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetUser(transaction *sqlx.Tx, param *dto.UserAddParam) (user *entity.User, err error) {
	query := "SELECT id_user, fk_role, name, email, password FROM user WHERE email = ?"

	row := transaction.QueryRowx(query, param.Email)

	if row == nil {
		return nil, errors.New("User not found")
	}

	err = row.Scan(user)

	return
}

func AddUser(transaction *sqlx.Tx, param *dto.UserAddParam) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (`%s`, `%s`, `%s`, `%s`)" +
			"VALUES (:fk_role, :name, :email, :password)",
		"user",
		"fk_role",
		"name",
		"email",
		"password",
	)
	if _, err := transaction.NamedExec(query, param); err != nil {
		return err
	}
	return nil
}