package repository

import (
	"fmt"
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Info(id int) (model.UserFull, error) {
	var user model.UserFull
	query := fmt.Sprintf("SELECT id, phone, name, role FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&user, query, id)
	if err != nil {
		fmt.Println(err)
		return model.UserFull{}, err
	}

	return user, nil
}

func (r *UserPostgres) SetRole(id int, role int) (int, error) {
	query := fmt.Sprintf("UPDATE users SET Role = $1 WHERE Id = $2")
	_, err := r.db.Exec(query, role, id)

	return 0, err
}
