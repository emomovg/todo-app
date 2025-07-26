package models

type User struct {
	Id       int    `json:"-"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
