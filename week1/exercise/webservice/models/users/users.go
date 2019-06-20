package users

import (
	"webservice/db"
	"webservice/dto"
)

func GetListUser() (users []dto.User, err error) {
	err = db.GetConn().Select(&users, "SELECT * FROM users")
	return
}

func GetUserById(id string) (user dto.User, err error) {
	err = db.GetConn().Get(&user, "SELECT * FROM users WHERE id = ?", id)
	return
}

func AddUser(user *dto.User) (err error) {
	transaction := db.GetConn().MustBegin()
	result, err := transaction.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	if err != nil {
		transaction.Rollback()
		return
	}
	userId, err := result.LastInsertId()
	if err != nil {
		transaction.Rollback()
		return
	}
	user.Id = userId
	transaction.Commit()
	return nil
}

func UpdateUser(user dto.User) (err error) {
	transaction := db.GetConn().MustBegin()
	transaction.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, user.Id)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return nil
}
