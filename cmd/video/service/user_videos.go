package service

import (
	"context"

	"github.com/YuleZhang/douyin/cmd/video/dao/db"
	"github.com/YuleZhang/douyin/kitex_gen/video"

	"github.com/YuleZhang/douyin/cmd/video/pack"
)

type DouyinPublishListService struct {
	ctx context.Context
}

func NewDouyinPublishListService(ctx context.Context) *DouyinPublishListService {
	return &DouyinPublishListService{ctx: ctx}
}

func (s *DouyinPublishListService) PublishList(req *video.DouyinPublishListRequest) ([]*video.Video, error) {
	modelVideos, err := db.QueryUserVideoList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.Videos(s.ctx, modelVideos), nil
}
