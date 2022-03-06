package service

import "github.com/heroku/go-getting-started/pkg/repository"

type AnswerService struct {
	repo repository.Answer
}

func NewAnswerService(repo repository.Answer) *AnswerService {
	return &AnswerService{repo: repo}
}

func (s *AnswerService) CreateAnswer(text string) (int, error) {
	return s.repo.Create(text)
}
