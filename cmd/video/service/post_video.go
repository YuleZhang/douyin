package service

import (
	"context"

	"github.com/yulezhang/douyin/cmd/video/dao/db"
	"github.com/yulezhang/douyin/kitex_gen/video"
)

type DouyinPublishActionService struct {
	ctx context.Context
}

func NewDouyinPublishActionService(ctx context.Context) *DouyinPublishActionService {
	return &DouyinPublishActionService{ctx: ctx}
}

func (s *DouyinPublishActionService) PublishAction(req *video.DouyinPublishActionRequest) error {
	err := db.CreateVideo(s.ctx, db.Video{
		UserID:        req.Video.Author.Id,
		PlayUrl:       req.Video.PlayUrl,
		CoverUrl:      req.Video.CoverUrl,
		FavoriteCount: req.Video.FavoriteCount,
		CommentCount:  req.Video.CommentCount,
		IsFavorite:    req.Video.IsFavorite,
		Title:         req.Video.Title,
	})
	if err != nil {
		return err
	}
	return nil
}
