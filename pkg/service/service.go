package service

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/pkg/repository"
)

type Answer interface {
	CreateAnswer(text string) (int, error)
}

type User interface {
	Info(id int) (model.UserFull, error)
	SetRole(id int, role int) (int, error)
}

type Authentication interface {
	CheckAuth(phone string) (int, error)
	CreateUser(user model.SignUpInput) (int, error)
	GenerateToken(name, phone string) (string, error)
	ParseToken(token string) (int, error)
}

type Company interface {
	GetCompany(company model.Company) (int, error)
}

type Group interface {
	CreateGroup(title string, userId int) (int, error)
	AddUserInGroup(groupAdd model.GroupAddUserInput) (int, error)
	GetAllGroupUser(userId int) ([]model.Group, error)
}

type Service struct {
	Answer
	Authentication
	Company
	Group
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Answer:         NewAnswerService(repos),
		Authentication: NewAuthService(repos),
		Company:        NewCompanyService(repos),
		Group:          NewGroupService(repos),
		User:           NewUserService(repos),
	}
}
