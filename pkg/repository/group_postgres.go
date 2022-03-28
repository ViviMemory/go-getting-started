package repository

import (
	"fmt"
	"github.com/heroku/go-getting-started/model"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
)

type GroupPostgres struct {
	db *sqlx.DB
}

func NewGroupPostgres(db *sqlx.DB) *GroupPostgres {
	return &GroupPostgres{db: db}
}

func (r *GroupPostgres) DetailGroup(groupId int) ([]model.UserGroup, error) {
	var users = []model.UserGroup{}
	query := fmt.Sprintf("SELECT users.id, users.username, users.phone FROM users INNER JOIN group_user ON group_user.user_id=users.id and group_user.group_company_id=$1 and group_user.status=3")
	err := r.db.Select(&users, query, groupId)

	if err != nil {
		fmt.Println(err)
		return users, err
	}
	return users, nil
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func (r *GroupPostgres) InviteUserInGroup(groupId int, phone string) (int, error) {
	var userId int
	query := fmt.Sprintf("select id from %s where phone=$1", userTable)
	err := r.db.Get(&userId, query, phone)

	if err != nil {
		fmt.Println(err)
		return userId, err
	}

	var othersIds []int
	query = fmt.Sprintf("select group_user.group_company_id from group_user where user_id=$1")
	err = r.db.Select(&othersIds, query, userId)

	if err != nil {
		fmt.Println(err)
		return userId, err
	}

	if itemExists(othersIds, groupId) {
		return 0, nil
	}

	query = fmt.Sprintf("INSERT INTO group_user (group_company_id, user_id, status) values ($1, $2, 1) RETURNING id")

	var idInsert int

	row := r.db.QueryRow(query, groupId, userId)
	if err := row.Scan(&idInsert); err != nil {
		return 0, err
	}

	return idInsert, nil
}

func (r *GroupPostgres) ListInviteUserInGroup(userId int) ([]model.GroupList, error) {
	var groups = []model.GroupList{}
	query := fmt.Sprintf("select group_company.id, group_company.title, group_user.status from group_company inner join group_user on group_user.group_company_id=group_company.id and not group_user.status=4 and group_user.user_id=$1")
	err := r.db.Select(&groups, query, userId)

	if err != nil {
		fmt.Println(err)
		return groups, err
	}

	return groups, nil
}

func (r *GroupPostgres) ActiveInviteUserInGroup(userId int, groupId int, isReject bool) error {
	query := fmt.Sprintf("update group_user set status=$1 where user_id=$2 and group_company_id=$3")
	var act int
	if isReject {
		act = 2
	} else {
		act = 3
	}
	_, err := r.db.Exec(query, act, userId, groupId)
	return err
}

func (r *GroupPostgres) CreateGroup(title string, userId int) (int, error) {
	if err := r.createGroupTable(); err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("insert into group_company (title, company_id) values ($1, 1) returning id")
	//query := fmt.Sprintf("INSERT INTO %s (title, main_user_id) values ($1, $2) RETURNING id", groupTable)

	row := r.db.QueryRow(query, title)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	var idQuery int
	query = fmt.Sprintf("insert into group_user (group_company_id, user_id, status) values ($1, $2, $3) returning id")
	//query := fmt.Sprintf("INSERT INTO %s (title, main_user_id) values ($1, $2) RETURNING id", groupTable)

	row = r.db.QueryRow(query, id, userId, 4)
	if err := row.Scan(&idQuery); err != nil {
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
