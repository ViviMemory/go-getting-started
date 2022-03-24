package repository

import (
	"fmt"
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
	"time"
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

func (r *TestPostgres) SaveResultTest(userId int, testId int, percentRight int) (int, error) {
	var id int
	currentTime := time.Now()
	query := fmt.Sprintf("insert into test_history (users_id, tests_id, percent_right, datetime_test) values($1, $2, $3, $4) returning id")
	row := r.db.QueryRow(query, userId, testId, percentRight, currentTime.Format("2006.01.02 15:04:05"))
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TestPostgres) HistoryMyTests(userId int) ([]model.TestHistoryItem, error) {
	var tests []model.TestHistoryItem
	query := fmt.Sprintf(" SELECT DISTINCT tests.title as title,test_history.datetime_test as datetime, test_history.percent_right as percent_right FROM test_history INNER JOIN tests ON tests.id=test_history.tests_id and test_history.users_id=$1")
	err := r.db.Select(&tests, query, userId)
	if err != nil {
		return tests, err
	}
	return tests, nil
}

func (r *TestPostgres) HistoryAllTests() ([]model.TestHistoryItem, error) {
	var tests []model.TestHistoryItem
	query := fmt.Sprintf(" SELECT DISTINCT tests.title as title,test_history.datetime_test as datetime, test_history.percent_right as percent_right FROM test_history INNER JOIN tests ON tests.id=test_history.tests_id")
	err := r.db.Select(&tests, query)
	if err != nil {
		return tests, err
	}
	return tests, nil
}

func (r *TestPostgres) DetailTest(testId int) (model.TestDetailOutput, error) {
	var result = model.TestDetailOutput{}

	//query := fmt.Sprintf("select title from tests where tests.tests_id=$1")
	//err := r.db.Select(&questionsId, query, testId)

	var questionsId []model.QuestionsTimeOutput
	// get all questions
	query := fmt.Sprintf("select distinct test_question.id, test_question.title from test_question where test_question.tests_id=$1")
	err := r.db.Select(&questionsId, query, testId)
	if err != nil {
		return result, err
	}

	// get all answers for all question
	for _, item := range questionsId {
		var answers []model.AnswerOutput
		query = fmt.Sprintf("select distinct question_answers.id, question_answers.title, question_answers.is_right from question_answers where question_answers.test_question_id=$1")
		err := r.db.Select(&answers, query, item.Id)
		if err != nil {
			return result, err
		}
		result.Questions = append(result.Questions, model.QuestionsOutput{
			Id:      item.Id,
			Title:   item.Title,
			Answers: answers,
		})
	}

	return result, nil
}
