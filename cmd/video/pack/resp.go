package pack

import (
	"errors"

	"github.com/yulezhang/douyin/kitex_gen/user"
	"github.com/yulezhang/douyin/pkg/errno"
)

func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
