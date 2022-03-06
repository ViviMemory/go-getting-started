package repository

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
)

type Answer interface {
	Create(text string) (int, error)
}

type Authentication interface {
	CheckAuth(phone string) (int, error)
	CreateUser(user model.SignUpInput) (int, error)
	GetUser(name, phone string) (model.User, error)
}

type Company interface {
	CreateCompany(company model.Company) (int, error)
	GetCompany(company model.Company) (int, error)
}

type Group interface {
	CreateGroup(title string, userId int) (int, error)
	AddUserInGroup(groupAdd model.GroupAddUserInput) (int, error)
	GetAllGroupUser(userId int) ([]model.Group, error)
}

type Repository struct {
	Answer
	Authentication
	Company
	Group
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Answer:         NewAnswerPostgres(db),
		Authentication: NewAuthPostgres(db),
		Company:        NewCompanyPostgres(db),
		Group:          NewGroupPostgres(db),
	}
}
