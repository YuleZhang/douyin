package service

import (
	"context"

	"github.com/YuleZhang/douyin/cmd/video/dao/db"
	"github.com/YuleZhang/douyin/kitex_gen/video"
)

type DouyinPublishActionService struct {
	ctx context.Context
}

func NewDouyinPublishActionService(ctx context.Context) *DouyinPublishActionService {
	return &DouyinPublishActionService{ctx: ctx}
}

func (s *DouyinPublishActionService) PublishAction(req *video.DouyinPublishActionRequest) error {
	err := db.CreateVideo(s.ctx, db.Video{
		UserID:        req.UserId,
		PlayUrl:       req.FileUrl,
		CoverUrl:      req.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         req.Title,
	})
	if err != nil {
		return err
	}
	return nil
}
