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
	DetailTest(testId int) (model.TestDetailOutput, error)
	SaveResultTest(userId int, testId int, percentRight int) (int, error)
	HistoryMyTests(userId int) ([]model.TestHistoryItem, error)
	HistoryAllTests() ([]model.TestHistoryAllItem, error)
}

type Question interface {
	AddQuestionInTest(title string, testId int) (int, error)
	AddAnswerInQuestion(testId int, title string, isRight bool) error
}

type Service struct {
	Answer
	Authentication
	Company
	Group
	User
	TestCategory
	Test
	Question
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Answer:         NewAnswerService(repos),
		Authentication: NewAuthService(repos),
		Company:        NewCompanyService(repos),
		Group:          NewGroupService(repos),
		User:           NewUserService(repos),
		TestCategory:   NewTestCategoryService(repos),
		Test:           NewTestService(repos),
		Question:       NewQuestionService(repos),
	}
}
