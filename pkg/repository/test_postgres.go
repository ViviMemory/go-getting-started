package repository

import (
	"fmt"
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
)

type TestPostgres struct {
	db *sqlx.DB
}

func NewTestPostgres(db *sqlx.DB) *TestPostgres {
	return &TestPostgres{db: db}
}

func (r *TestPostgres) CreatedTest(title string, categoryId int, accessPrivate bool, userId int) (int, error) {
	var id int
	query := fmt.Sprintf("insert into tests (title, test_category_id, access_private, users_id) values($1, $2, $3, $4) returning id")
	row := r.db.QueryRow(query, title, categoryId, accessPrivate, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TestPostgres) AddPrivateTestInGroup(testId int, groupId int) (int, error) {
	var id int
	query := fmt.Sprintf("insert into tests_group (tests_id, group_company_id) values($1, $2) returning id")
	row := r.db.QueryRow(query, testId, groupId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TestPostgres) AllTest(userId int) (model.TestOutput, error) {
	var result = model.TestOutput{}

	var privateTests = []model.TestOfUser{}
	query := fmt.Sprintf("SELECT DISTINCT tests.id, tests.title, group_company.title as group_title FROM tests INNER JOIN tests_group ON tests.id=tests_group.tests_id INNER JOIN group_user ON group_user.user_id=$1 AND group_user.group_company_id=tests_group.group_company_id INNER JOIN group_company ON group_company.id=tests_group.group_company_id")
	err := r.db.Select(&privateTests, query, userId)

	if err != nil {
		return result, err
	}

	result.Private = privateTests

	var publicTests = []model.TestPublic{}
	query = fmt.Sprintf("select id, title from tests where access_private=false")
	err = r.db.Select(&publicTests, query)

	if err != nil {
		return result, err
	}

	result.Publish = publicTests

	return result, err
}
