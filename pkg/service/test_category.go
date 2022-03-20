package service

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/pkg/repository"
)

type TestCategoryService struct {
	repo repository.TestCategory
}

func NewTestCategoryService(repo repository.TestCategory) *TestCategoryService {
	return &TestCategoryService{repo: repo}
}

func (s *TestCategoryService) CategoriesList() ([]model.CategoryInfo, error) {
	return s.repo.CategoriesList()
}

func (s *TestCategoryService) AddCategoryTest(title string) (int, error) {
	return s.repo.AddCategoryTest(title)
}
