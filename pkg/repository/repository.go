package repository

import "github.com/jmoiron/sqlx"

type Answer interface {
	Create(text string) (int, error)
}

type Authentication interface {
	CheckAuth(phone string) (int, error)
}

type Repository struct {
	Answer
	Authentication
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Answer:         NewAnswerPostgres(db),
		Authentication: NewAuthPostgres(db),
	}
}
