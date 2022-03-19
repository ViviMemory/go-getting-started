package service

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Info(id int) (model.UserFull, error) {
	return s.repo.Info(id)
}

func (s *UserService) SetRole(id int, role int) (int, error) {
	return s.repo.SetRole(id, role)
}
