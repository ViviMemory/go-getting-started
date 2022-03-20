package repository

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
)

type Answer interface {
	Create(text string) (int, error)
}

type User interface {
	Info(id int) (model.UserFull, error)
	SetRole(id int, role int) (int, error)
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
	DetailGroup(groupId int) ([]model.UserGroup, error)
	InviteUserInGroup(groupId int, phone string) (int, error)
	ListInviteUserInGroup(userId int) ([]model.GroupList, error)
	ActiveInviteUserInGroup(userId int, groupId int, isReject bool) error
}

type TestCategory interface {
	CategoriesList() ([]model.CategoryInfo, error)
	AddCategoryTest(title string) (int, error)
}

type Test interface {
	CreatedTest(title string, categoryId int, accessPrivate bool, userId int) (int, error)
	AddPrivateTestInGroup(testId int, groupId int) (int, error)
	AllTest(userId int) (model.TestOutput, error)
}

type Repository struct {
	Answer
	Authentication
	Company
	Group
	User
	TestCategory
	Test
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Answer:         NewAnswerPostgres(db),
		Authentication: NewAuthPostgres(db),
		Company:        NewCompanyPostgres(db),
		Group:          NewGroupPostgres(db),
		User:           NewUserPostgres(db),
		TestCategory:   NewTestCategoryPostgres(db),
		Test:           NewTestPostgres(db),
	}
}
