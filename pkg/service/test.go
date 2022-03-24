package service

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/pkg/repository"
)

type TestService struct {
	repo repository.Test
}

func (s *TestService) HistoryMyTests(userId int) ([]model.TestHistoryItem, error) {
	return s.repo.HistoryMyTests(userId)
}

func (s *TestService) HistoryAllTests() ([]model.TestHistoryAllItem, error) {
	return s.repo.HistoryAllTests()
}

func (s *TestService) DetailTest(testId int) (model.TestDetailOutput, error) {
	return s.repo.DetailTest(testId)
}

func NewTestService(repo repository.Test) *TestService {
	return &TestService{repo: repo}
}

func (s *TestService) SaveResultTest(userId int, testId int, percentRight int) (int, error) {
	return s.repo.SaveResultTest(userId, testId, percentRight)
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
