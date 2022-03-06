package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AnswerPostgres struct {
	db *sqlx.DB
}

func NewAnswerPostgres(db *sqlx.DB) *AnswerPostgres {
	return &AnswerPostgres{db: db}
}
func (r *AnswerPostgres) Create(text string) (int, error) {
	tx, err := r.db.Beginx()

	if err != nil {
		return 0, err
	}

	if err := r.createTable(); err != nil {
		log.Printf("it is not possible to create a table of answer options")
		return 0, err
	}

	var id int
	insertQuery := fmt.Sprintf("INSERT INTO answer (text) VALUES ($1) RETURNING id")
	row := tx.QueryRow(insertQuery, text)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, nil
}

func (r *AnswerPostgres) createTable() error {
	createAnswerTable := "CREATE TABLE IF NOT EXISTS answer (" +
		"id serial not null  unique, " +
		"text varchar(255) not null"

	if _, err := r.db.Exec(createAnswerTable); err != nil {
		return err
	}

	return nil
}
