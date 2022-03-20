package model

type Group struct {
	Id         int    `json:"-" db:"id"`
	Name       string `json:"title" db:"title" binding:"required"`
	MainUserId string `json:"main_user_id" db:"main_user_id" binding:"required"`
}

type GroupInput struct {
	Name string `json:"name" binding:"required"`
}

type GroupId struct {
	Id       int  `json:"group_id"`
	IsReject bool `json:"is_reject"`
}

type GroupAddUserInput struct {
	GroupId int    `json:"group_id" `
	Phone   string `json:"phone" `
}

type GroupList struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

type GroupDeleteUserInput struct {
	GroupId int    `json:"group_id" `
	Phone   string `json:"phone" `
}

type GroupAll struct {
	Title string `json:"title" db:"title" binding:"required"`
	Users string `json:"user" db:"users" binding:"required"`
}

type GroupDetail struct {
	Users []UserGroup `json:"users"`
}

type UserGroup struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"username" db:"username" binding:"required"`
	Phone string `json:"phone" db:"phone" binding:"required"`
}
