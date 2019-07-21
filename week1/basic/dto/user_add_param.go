package dto

type UserAddParam struct {
	FkRole 		int		`json:"role_id"`
	Name 		string
	Email		string
	Password 	string
}