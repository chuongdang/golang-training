package entity

type User struct {
	FkRole   int `db:"fk_role"`
	UserCommon
}

type UserCommon struct {
	Name     string
	Email    string
	Password string
}