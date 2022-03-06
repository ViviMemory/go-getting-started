package model

type User struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
