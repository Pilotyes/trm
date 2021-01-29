package store

import "trm/internal/model"

const (
	//UserTypeR R&D ...
	UserTypeR = iota
	//UserTypeC Constructors ...
	UserTypeC
)

//UserList ...
var UserList = map[string]*model.User{
	"user1": &model.User{
		ID:       int64(1),
		Login:    "user1",
		Password: "pass1",
		UserType: UserTypeR,
	},
	"user2": &model.User{
		ID:       int64(2),
		Login:    "user2",
		Password: "pass2",
		UserType: UserTypeC,
	},
}

//FindUser ...
func FindUser(userName string) *model.User {
	if u, ok := UserList[userName]; ok {
		return u
	}

	return nil
}
