package service

import (
	"context"

	"github.com/RaymondCode/simple-demo/cmd/video/dao/db"
	"github.com/RaymondCode/simple-demo/kitex_gen/video"

	"github.com/RaymondCode/simple-demo/cmd/video/pack"
)

type DouyinFeedService struct {
	ctx context.Context
}

func NewDouyinFeedService(ctx context.Context) *DouyinFeedService {
	return &DouyinFeedService{ctx: ctx}
}

func (s *DouyinFeedService) DouyinFeed(req *video.DouyinFeedRequest) ([]*video.Video, error) {
	modelVideos, err := db.QueryVideoFeed(s.ctx)
	if err != nil {
		return nil, err
	}
	return pack.Videos(s.ctx, modelVideos), nil
}
