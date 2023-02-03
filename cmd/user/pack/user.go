package pack

import (
	"tiktok/cmd/user/dal/db"
	"tiktok/kitex_gen/user"
)

// User pack user info
/*
is_follow需要另外获取
*/
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{Id: int64(u.UserId)}
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
