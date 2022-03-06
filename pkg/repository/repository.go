package repository

type Answer interface {
	Create(text string) (int, error)
}

type Repository struct {
	Answer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Answer: NewAnswerPostgres(db),
	}
}
