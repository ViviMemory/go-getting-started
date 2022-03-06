package service

import "github.com/heroku/go-getting-started/pkg/repository"

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repo repository.Authentication) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CheckAuth(phone string) (int, error) {
	id, err := s.repo.CheckAuth(phone)
	return id, err
}
