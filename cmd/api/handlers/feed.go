package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/YuleZhang/douyin/cmd/api/rpc"
	"github.com/YuleZhang/douyin/kitex_gen/video"
	"github.com/YuleZhang/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []*video.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

func DouyinFeed(c *gin.Context) {
	resp, err := rpc.DouyinFeed(context.Background(), &video.DouyinFeedRequest{})
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: Err.ErrCode, StatusMsg: Err.ErrMsg},
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: errno.SuccessCode,
			StatusMsg:  resp.BaseResp.StatusMsg,
		},
		VideoList: resp.VideoList,
		NextTime:  time.Now().Unix(),
	})
}
