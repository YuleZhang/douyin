package service

import (
	"context"

	"github.com/YuleZhang/douyin/kitex_gen/user"

	"github.com/YuleZhang/douyin/cmd/user/dao/db"

	"github.com/YuleZhang/douyin/cmd/user/pack"
)

type DouyinUserService struct {
	ctx context.Context
}

func NewDouyinUserService(ctx context.Context) *DouyinUserService {
	return &DouyinUserService{ctx: ctx}
}

// 暂时只支持查单用户
func (s *DouyinUserService) UserInfo(req *user.DouyinUserRequest) ([]*user.User, error) {
	modelUsers, err := db.QueryUserWithID(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
