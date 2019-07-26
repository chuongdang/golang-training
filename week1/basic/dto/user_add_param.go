package dto

import "basic/entity"

type UserAddParam struct {
	FkRole 		int		`json:"role_id" db:"fk_role"`
	entity.UserCommon
}