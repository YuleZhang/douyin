package pack

import (
	"github.com/YuleZhang/douyin/cmd/user/dao/db"
	"github.com/YuleZhang/douyin/kitex_gen/user"
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{Id: int64(u.ID), Name: u.Name, FollowCount: &u.FollowCount, FollowerCount: &u.FollowerCount, IsFollow: u.IsFollow}
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
