package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/yulezhang/douyin/pkg/errno"

	"github.com/yulezhang/douyin/kitex_gen/user"

	"github.com/yulezhang/douyin/cmd/user/dao/db"
)

type DouyinUserLoginService struct {
	ctx context.Context
}

func NewDouyinUserLoginService(ctx context.Context) *DouyinUserLoginService {
	return &DouyinUserLoginService{
		ctx: ctx,
	}
}

func (s *DouyinUserLoginService) UserLogin(req *user.DouyinUserLoginRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	users, err := db.QueryUserWithName(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(u.ID), nil
}
