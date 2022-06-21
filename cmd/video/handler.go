package main

import (
	"context"

	"github.com/yulezhang/douyin/cmd/video/pack"
	"github.com/yulezhang/douyin/cmd/video/service"
	"github.com/yulezhang/douyin/kitex_gen/video"
	"github.com/yulezhang/douyin/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// DouyinFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DouyinFeed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	resp = new(video.DouyinFeedResponse)
	video_list, err := service.NewDouyinFeedService(ctx).DouyinFeed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = video_list
	return resp, nil
}

// DouyinPublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DouyinPublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	resp = new(video.DouyinPublishActionResponse)
	err = service.NewDouyinPublishActionService(ctx).PublishAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DouyinPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DouyinPublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	resp = new(video.DouyinPublishListResponse)
	video_list, err := service.NewDouyinPublishListService(ctx).PublishList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = video_list
	return resp, nil
}

// DouyinFavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DouyinFavoriteList(ctx context.Context, req *video.DouyinFavoriteListRequest) (resp *video.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}
