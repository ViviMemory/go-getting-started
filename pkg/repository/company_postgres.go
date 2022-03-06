package repository

import (
	"fmt"
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
)

type CompanyPostgres struct {
	db *sqlx.DB
}

func NewCompanyPostgres(db *sqlx.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db}
}

func (r *CompanyPostgres) CreateCompany(company model.Company) (int, error) {
	if err := r.createCompanyTable(); err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, address) values ($1, $2) RETURNING id", companyTable)

	row := r.db.QueryRow(query, company.Title, company.Address)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CompanyPostgres) GetCompany(company model.Company) (int, error) {
	if err := r.createCompanyTable(); err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE title=$1 AND address=$2", companyTable)
	err := r.db.Get(&id, query, company.Title, company.Address)

	if err != nil || id == 0 {
		return r.CreateCompany(company)
	}

	return id, err
}

func (r *CompanyPostgres) createCompanyTable() error {
	createCompanyTable := "CREATE TABLE IF NOT EXISTS company (" +
		"id serial not null unique," +
		"title varchar(255) not null," +
		"address varchar(255) not null" +
		")"

	if _, err := r.db.Exec(createCompanyTable); err != nil {
		return err
	}
	return nil
}
