package service

import "github.com/heroku/go-getting-started/pkg/repository"

type QuestionService struct {
	repo repository.Question
}

func NewQuestionService(repo repository.Question) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) AddQuestionInTest(title string, testId int) (int, error) {
	return s.repo.AddQuestionInTest(title, testId)
}

func (s *QuestionService) AddAnswerInQuestion(testId int, title string, isRight bool) error {
	return s.repo.AddAnswerInQuestion(testId, title, isRight)
}
