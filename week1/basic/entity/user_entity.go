package entity

type User struct {
	ID       int `db:"id_user" json:"id"`
	FkRole   int `db:"fk_role" json:"role_id"`
	UserCommon
}

type UserCommon struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}