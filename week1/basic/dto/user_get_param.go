package dto

type UserGetParam struct {
	Limit  int `form:"limit"`
	Page   int `form:"page"`
}
