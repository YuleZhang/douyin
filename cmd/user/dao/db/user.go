package db

import (
	"context"
	"fmt"

	"github.com/yulezhang/douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// DouyinUser get list of user info according to user_id list
func QueryUserWithID(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Debug().Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// DouyinUserRegister create user info
func UserRegister(ctx context.Context, user User) (int64, error) {
	err := DB.WithContext(ctx).Create(&user).Error
	return int64(user.ID), err
}

// QueryUser query list of user info according to username
func QueryUserWithName(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	fmt.Println("username to be search", userName)
	if err := DB.WithContext(ctx).Debug().Where("name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	fmt.Println("search result", res)
	return res, nil
}
