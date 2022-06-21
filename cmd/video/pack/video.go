package pack

import (
	"context"

	"github.com/yulezhang/douyin/cmd/video/dao/db"
	"github.com/yulezhang/douyin/kitex_gen/user"
	"github.com/yulezhang/douyin/kitex_gen/video"
)

// User pack user info
func Video(ctx context.Context, v *db.Video) *video.Video {
	if v == nil {
		return nil
	}
	u, err := db.QueryUserWithVideo(ctx, v.UserID)
	if err != nil {
		return &video.Video{}
	}
	return &video.Video{
		Id:            int64(v.ID),
		PlayUrl:       v.PlayUrl,
		Author:        &user.User{Id: int64(u.Id), Name: u.Name, FollowCount: u.FollowCount, FollowerCount: u.FollowerCount, IsFollow: u.IsFollow},
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

// Users pack list of user info
func Videos(ctx context.Context, us []*db.Video) []*video.Video {
	videos := make([]*video.Video, 0)
	for _, u := range us {
		if video2 := Video(ctx, u); video2 != nil {
			videos = append(videos, video2)
		}
	}
	return videos
}
