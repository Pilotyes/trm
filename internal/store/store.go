package store

import "trm/internal/model"

const (
	//UserTypeM - пользователь группы Maitenece
	UserTypeM = iota
	//UserTypeC Contractor - пользователь группы контрактников
	UserTypeC
)

//UserList - БД пользователей, зарегистрированных в системе
var UserList = map[string]*model.User{
	"user1": {
		ID:       int64(1),
		Login:    "user1",
		Password: "pass1",
		UserType: UserTypeM,
	},
	"user2": {
		ID:       int64(2),
		Login:    "user2",
		Password: "pass2",
		UserType: UserTypeC,
	},
}

//FindUser - поиск в БД зарегистрированных пользователей по полям структуры User
func FindUser(userName string) *model.User {
	if u, ok := UserList[userName]; ok {
		return u
	}

	return nil
}
