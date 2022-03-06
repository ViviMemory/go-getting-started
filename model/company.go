package model

type Company struct {
	Title   string `json:"title" binding:"required"`
	Address string `json:"address" binding:"required"`
}
