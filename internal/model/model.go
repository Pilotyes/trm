package model

//User - структура пользователя
type User struct {
	ID       int64
	Login    string
	Password string
	UserType int
}
