package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/YuleZhang/douyin/cmd/user/dao/db"
	"github.com/YuleZhang/douyin/kitex_gen/user"
	"github.com/YuleZhang/douyin/pkg/errno"
)

type DouyinUserRegisterService struct {
	ctx context.Context
}

func NewDouyinUserRegisterService(ctx context.Context) *DouyinUserRegisterService {
	return &DouyinUserRegisterService{ctx: ctx}
}

func (s *DouyinUserRegisterService) UserRegist(req *user.DouyinUserRegisterRequest) (int64, error) {
	user, err := db.QueryUserWithName(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(user) != 0 {
		return 0, errno.UserAlreadyExistErr
	}
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	uid, err := db.UserRegister(s.ctx, db.User{
		Name:          req.Username,
		Password:      passWord,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	})
	return uid, err
}
