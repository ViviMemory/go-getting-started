package model

type TestHistoryItem struct {
	Title        string `json:"title" db:"title"`
	PercentRight int    `json:"percent_right" db:"percent_right"`
	Datetime     string `json:"datetime"`
}

type TestHistoryAllItem struct {
	Title        string `json:"title" db:"title"`
	PercentRight int    `json:"percent_right" db:"percent_right"`
	Datetime     string `json:"datetime"`
	UserPhone    string `json:"user_phone" db:"user_phone"`
	UserName     string `json:"user_name" db:"user_name"`
}

type TestInput struct {
	Title         string `json:"title" db:"title"`
	AccessPrivate bool   `json:"access_private"`
	//CategoryId    int    `json:"category_id"`
	GroupId   int              `json:"group_id"`
	Questions []QuestionsInput `json:"questions"`
}

type TestDetailOutput struct {
	Title     string            `json:"title" db:"title"`
	Questions []QuestionsOutput `json:"questions"`
}

type QuestionsInput struct {
	Title   string        `json:"title"`
	Answers []AnswerInput `json:"answers"`
}

type QuestionsTimeOutput struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type QuestionsOutput struct {
	Id      int            `json:"id"`
	Title   string         `json:"title"`
	Answers []AnswerOutput `json:"answers"`
}

type AnswerInput struct {
	Title   string `json:"title"`
	IsRight bool   `json:"is_right"`
}

type AnswerOutput struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	IsRight bool   `json:"is_right" db:"is_right"`
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
