package model

type CategoryInfo struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

type CategoryCreated struct {
	Title string `json:"title" db:"title"`
}
