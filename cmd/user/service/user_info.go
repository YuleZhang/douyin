package service

import (
	"context"

	"github.com/RaymondCode/simple-demo/kitex_gen/user"

	"github.com/RaymondCode/simple-demo/cmd/user/dao/db"

	"github.com/RaymondCode/simple-demo/cmd/user/pack"
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
