package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type QuestionPostgres struct {
	db *sqlx.DB
}

func NewQuestionPostgres(db *sqlx.DB) *QuestionPostgres {
	return &QuestionPostgres{
		db: db,
	}
}

func (r *QuestionPostgres) AddQuestionInTest(title string, testId int) (int, error) {
	var id int
	query := fmt.Sprintf("insert into test_question (title, tests_id) values($1, $2) returning id")
	row := r.db.QueryRow(query, title, testId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *QuestionPostgres) AddAnswerInQuestion(testId int, title string, isRight bool) error {
	var id int
	query := fmt.Sprintf("insert into question_answers (title, test_question_id, is_right) values($1, $2, $3) returning id")
	row := r.db.QueryRow(query, title, testId, isRight)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}
