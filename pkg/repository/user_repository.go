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
	var user model.UserTable
	query := fmt.Sprintf("SELECT id, username, phone, company_id, role_id FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&user, query, id)
	if err != nil {
		fmt.Println(err)
		return model.UserFull{}, err
	}

	var company model.CompanyFull

	if user.CompanyId != 0 {
		query = fmt.Sprintf("SELECT title, address FROM company where id=$1")
		err = r.db.Get(&company, query, user.CompanyId)

		if err != nil {
			fmt.Println(err)
			return model.UserFull{}, err
		}

		var groups []model.GroupCompanyInInfo
		query := fmt.Sprintf("select group_company.id as id, group_company.title as title from group_company inner join group_user on group_user.user_id=$1 and group_user.group_company_id=group_company.id")
		//query = fmt.Sprintf("SELECT id, title FROM group_company where company_id=$1")
		err = r.db.Select(&groups, query, user.CompanyId)

		if err != nil {
			fmt.Println(user.CompanyId)
			fmt.Println(err)
			return model.UserFull{}, err
		}

		company.Groups = groups
	}

	return model.UserFull{
		Id:          user.Id,
		Name:        user.Name,
		Phone:       user.Phone,
		CompanyUser: company,
		Role:        user.Role,
	}, nil
}

func (r *UserPostgres) SetRole(id int, role int) (int, error) {
	query := fmt.Sprintf("UPDATE users SET Role = $1 WHERE Id = $2")
	_, err := r.db.Exec(query, role, id)

	return 0, err
}
