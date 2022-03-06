package repository

import (
	"fmt"
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) GetUser(name, phone string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND phone=$2", userTable)
	err := r.db.Get(&user, query, name, phone)

	return user, err
}

func (r *AuthPostgres) CheckAuth(phone string) (int, error) {

	if err := r.createUserTable(); err != nil {
		return 0, err
	}

	var user model.User
	query := fmt.Sprintf("select id from %s where phone=$1", userTable)
	err := r.db.Get(&user, query, phone)

	return user.Id, err
}

func (r *AuthPostgres) CreateUser(user model.SignUpInput) (int, error) {

	if err := r.createUserTable(); err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, phone, role) values ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Name, user.Phone, user.RoleId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) createUserTable() error {
	createUserTable := "CREATE TABLE IF NOT EXISTS users (" +
		"id serial not null unique," +
		"name varchar(255) not null," +
		"phone varchar(255) not null," +
		"role int" +
		")"

	if _, err := r.db.Exec(createUserTable); err != nil {
		return err
	}
	return nil
}
