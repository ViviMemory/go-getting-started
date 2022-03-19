package model

type User struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type UserFull struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Role  string `json:"role_id" db:"role" binding:"required"`
}

/**
child - 0
parent - 1
*/
type SignUpInput struct {
	Phone  string `json:"phone" binding:"required"`
	Name   string `json:"name" binding:"required"`
	RoleId int    `json:"role_id" binding:"required"`
}

type SignInInput struct {
	Phone string `json:"phone" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
