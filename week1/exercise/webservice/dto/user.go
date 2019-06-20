package dto

type User struct {
	Id   int64  `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"username"`
}
