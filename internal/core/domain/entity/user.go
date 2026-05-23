package entity

type UserCtx struct {
	Id    string
	Email string
}

type User struct {
	Base
	Name     string
	Email    string
	Password string
}
