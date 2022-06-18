package db

import (
	"context"

	"github.com/RaymondCode/simple-demo/kitex_gen/user"
	"github.com/RaymondCode/simple-demo/pkg/constants"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserID        int64  `json:"user_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func QueryUserWithVideo(ctx context.Context, userID int64) (*user.User, error) {
	var u *user.User
	DB.WithContext(ctx).Debug().Table("user").Where("id = ?", userID).Find(&u)
	return u, nil
}

// 获取视频流
func QueryVideoFeed(ctx context.Context) ([]*Video, error) {
	res := make([]*Video, 0)
	DB.WithContext(ctx).Debug().Preload("User").Find(&res)
	return res, nil
}

// 查询用户视频列表
func QueryUserVideoList(ctx context.Context, userID int64) ([]*Video, error) {
	res := make([]*Video, 0)
	DB.WithContext(ctx).Debug().Preload("User").Where("ID = ?", userID).Find(&res)
	return res, nil
}

// 创建视频信息
func CreateVideo(ctx context.Context, video Video) error {
	return DB.WithContext(ctx).Debug().Create(video).Error
}
