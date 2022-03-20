package model

type TestInput struct {
	Title         string `json:"title" db:"title"`
	AccessPrivate bool   `json:"access_private"`
	CategoryId    int    `json:"category_id"`
	GroupId       int    `json:"group_id"`
}

type TestOutput struct {
	Publish []TestPublic `json:"publish"`
	Private []TestOfUser `json:"private"`
}

type TestOfUser struct {
	Id         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	GroupTitle string `json:"group_title" db:"group_title"`
}

type TestPublic struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}
