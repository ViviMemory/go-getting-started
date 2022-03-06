package service

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/pkg/repository"
)

type CompanyService struct {
	repo repository.Company
}

func NewCompanyService(repo repository.Company) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) GetCompany(company model.Company) (int, error) {
	return s.repo.GetCompany(company)
}
