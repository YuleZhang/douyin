package main

import (
	"context"

	"github.com/yulezhang/douyin/cmd/user/pack"
	"github.com/yulezhang/douyin/cmd/user/service"
	"github.com/yulezhang/douyin/kitex_gen/user"
	"github.com/yulezhang/douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DouyinUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp = new(user.DouyinUserResponse)

	// if len(req.UserIds) == 0 {
	// 	resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
	// 	return resp, nil
	// }
	user, err := service.NewDouyinUserService(ctx).UserInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = user
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DouyinUserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	uid, err := service.NewDouyinUserLoginService(ctx).UserLogin(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DouyinUserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) DouyinUserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewDouyinUserRegisterService(ctx).UserRegist(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
