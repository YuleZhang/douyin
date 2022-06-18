package rpc

import (
	"context"
	"time"

	"github.com/RaymondCode/simple-demo/kitex_gen/video"
	"github.com/RaymondCode/simple-demo/kitex_gen/video/videoservice"
	"github.com/RaymondCode/simple-demo/pkg/constants"
	"github.com/RaymondCode/simple-demo/pkg/errno"
	"github.com/RaymondCode/simple-demo/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
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
	videoClient = c
}

func DouyinFeed(ctx context.Context, req *video.DouyinFeedRequest) (*video.DouyinFeedResponse, error) {
	resp, err := videoClient.DouyinFeed(ctx, req)
	if err != nil {
		Err := errno.ConvertErr(err)
		return nil, errno.NewErrNo(Err.ErrCode, Err.ErrMsg)
	}

	return resp, nil
}

func DouyinPublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	resp, err := videoClient.DouyinPublishAction(ctx, req)
	if err != nil {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

func DouyinPublishList(ctx context.Context, req *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	resp, err := videoClient.DouyinPublishList(ctx, req)
	if err != nil {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}
