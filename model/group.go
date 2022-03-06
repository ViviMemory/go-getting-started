package model

type Group struct {
	Id         int    `json:"-" db:"id"`
	Name       string `json:"title" db:"title" binding:"required"`
	MainUserId string `json:"main_user_id" db:"main_user_id" binding:"required"`
}

type GroupInput struct {
	Name string `json:"name" binding:"required"`
}

type GroupAddUserInput struct {
	GroupId int    `json:"group_id" `
	Phone   string `json:"phone" `
}

type GroupDeleteUserInput struct {
	GroupId int    `json:"group_id" `
	Phone   string `json:"phone" `
}

type GroupAll struct {
	Title string `json:"title" db:"title" binding:"required"`
	Users string `json:"user" db:"users" binding:"required"`
}
