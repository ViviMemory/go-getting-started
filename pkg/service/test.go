package service

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/pkg/repository"
)

type TestService struct {
	repo repository.Test
}

func NewTestService(repo repository.Test) *TestService {
	return &TestService{repo: repo}
}

func (s *TestService) CreatedTest(title string, categoryId int, accessPrivate bool, userId int) (int, error) {
	return s.repo.CreatedTest(title, categoryId, accessPrivate, userId)
}

func (s *TestService) AddPrivateTestInGroup(testId int, groupId int) (int, error) {
	return s.repo.AddPrivateTestInGroup(testId, groupId)
}

func (s *TestService) AllTest(userId int) (model.TestOutput, error) {
	return s.repo.AllTest(userId)
}
