package model

type User struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type SignUpInput struct {
	Phone  string `json:"phone" binding:"required"`
	Name   string `json:"name" binding:"required"`
	RoleId int    `json:"role_id" binding:"required"`
}

type SignInInput struct {
	Phone string `json:"phone" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
