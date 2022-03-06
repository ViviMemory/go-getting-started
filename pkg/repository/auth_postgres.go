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

func (r *AuthPostgres) CheckAuth(phone string) (int, error) {

	if err := r.createUserTable(); err != nil {
		return 0, err
	}

	var user model.User
	query := fmt.Sprintf("select id from %s where name=$1", userTable)
	err := r.db.Get(&user, query, phone)

	return user.Id, err
}

func (r *AuthPostgres) createUserTable() error {
	createUserTable := "create table if not exists user (" +
		"id serial not null unique," +
		"name varchar(255) not null," +
		"phone varchar(255) not null," +
		"role int not null" +
		")"

	if _, err := r.db.Exec(createUserTable); err != nil {
		return err
	}
	return nil
}
