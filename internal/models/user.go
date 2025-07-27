package models

const TableName = "users"

type User struct {
	Id       int    `json:"-"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
