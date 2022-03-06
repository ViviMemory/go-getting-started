package repository

import (
	"fmt"
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
	"log"
)

type GroupPostgres struct {
	db *sqlx.DB
}

func NewGroupPostgres(db *sqlx.DB) *GroupPostgres {
	return &GroupPostgres{db: db}
}

func (r *GroupPostgres) CreateGroup(title string, userId int) (int, error) {
	if err := r.createGroupTable(); err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, main_user_id) values ($1, $2) RETURNING id", groupTable)

	row := r.db.QueryRow(query, title, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *GroupPostgres) AddUserInGroup(groupAdd model.GroupAddUserInput) (int, error) {

	if err := r.createGroupTable(); err != nil {
		return 0, err
	}
	if err := r.createUserGroupTable(); err != nil {
		return 0, err
	}

	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE phone=$1", userTable)
	err := r.db.Get(&user, query, groupAdd.Phone)

	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (group_id, user_id) VALUES ($1, $2) RETURNING id", userGroupTable)
	row := r.db.QueryRow(createListQuery, groupAdd.GroupId, user.Id)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *GroupPostgres) GetAllGroupUser(userId int) ([]model.Group, error) {
	if err := r.createGroupTable(); err != nil {
		return nil, err
	}

	query := "select id, title from groups where main_user_id=$1 "
	var lists []model.Group
	err := r.db.Select(&lists, query, userId)

	log.Print(lists)

	return lists, err
}

func (r *GroupPostgres) createGroupTable() error {
	createGroupTable := "CREATE TABLE IF NOT EXISTS groups (" +
		"id serial not null unique," +
		"title varchar(255) not null," +
		"main_user_id int references users (id) on delete cascade not null" +
		")"
	if _, err := r.db.Exec(createGroupTable); err != nil {
		return err
	}
	return nil
}

func (r *GroupPostgres) createUserGroupTable() error {
	createUserGroupTable := "CREATE TABLE IF NOT EXISTS user_groups (" +
		"id serial not null unique," +
		"group_id int references groups (id) on delete cascade not null," +
		"user_id int references users (id) on delete cascade not null" +
		")"
	if _, err := r.db.Exec(createUserGroupTable); err != nil {
		return err
	}
	return nil
}
