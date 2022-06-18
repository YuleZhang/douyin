package rpc

import (
	"context"
	"time"

	"github.com/RaymondCode/simple-demo/kitex_gen/user"
	"github.com/RaymondCode/simple-demo/kitex_gen/user/userservice"
	"github.com/RaymondCode/simple-demo/pkg/constants"
	"github.com/RaymondCode/simple-demo/pkg/errno"
	"github.com/RaymondCode/simple-demo/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// DouyinUserRegister create user info
func DouyinUserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (int64, error) {
	resp, err := userClient.DouyinUserRegister(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}

// DouyinUserLogin check user info
func DouyinUserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (int64, error) {
	resp, err := userClient.DouyinUserLogin(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}

func DouyinUserInfo(ctx context.Context, req *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	resp, err := userClient.DouyinUser(ctx, req)
	return resp, err
}
