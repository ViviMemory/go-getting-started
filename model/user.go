package model

type User struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" db:"username" binding:"required"`
	Phone string `json:"phone" db:"phone" binding:"required"`
}

type UserTable struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"username" db:"username" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	CompanyId int    `json:"company_id"  db:"company_id"`
	Role      int    `json:"role_id" db:"role_id" binding:"required"`
}

type UserFull struct {
	Id          int         `json:"-" db:"id"`
	Name        string      `json:"name" binding:"required"`
	Phone       string      `json:"phone" binding:"required"`
	CompanyUser CompanyFull `json:"company"`
	Role        int         `json:"role_id" db:"role" binding:"required"`
}

type GroupCompanyInInfo struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

type CompanyFull struct {
	Title   string               `json:"title"`
	Address string               `json:"address"`
	Groups  []GroupCompanyInInfo `json:"groups"`
}

type SignUpInput struct {
	Phone  string `json:"phone" binding:"required"`
	Name   string `json:"name" binding:"required"`
	RoleId int    `json:"role_id" binding:"required"`
}

type SignInInput struct {
	Phone string `json:"phone" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
