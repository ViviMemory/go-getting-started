package service

import "github.com/heroku/go-getting-started/pkg/repository"

type Answer interface {
	CreateAnswer(text string) (int, error)
}

type Service struct {
	Answer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Answer: NewAnswerService(repos),
	}
}
