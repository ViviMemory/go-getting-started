package repository

import (
	"fmt"
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
)

type TestCategoryPostgres struct {
	db *sqlx.DB
}

func NewTestCategoryPostgres(db *sqlx.DB) *TestCategoryPostgres {
	return &TestCategoryPostgres{db: db}
}

func (r *TestCategoryPostgres) CategoriesList() ([]model.CategoryInfo, error) {
	var categories = []model.CategoryInfo{}
	query := fmt.Sprintf("select id, title from test_category")
	err := r.db.Select(&categories, query)
	return categories, err
}

func (r *TestCategoryPostgres) AddCategoryTest(title string) (int, error) {
	var id int
	query := fmt.Sprintf("insert into test_category (title) values($1) returning id")
	row := r.db.QueryRow(query, title)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
