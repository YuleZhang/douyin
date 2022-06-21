package handlers

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/YuleZhang/douyin/cmd/api/rpc"
	"github.com/YuleZhang/douyin/kitex_gen/user"
	"github.com/YuleZhang/douyin/kitex_gen/video"
	"github.com/YuleZhang/douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

type UserVideoListResponse struct {
	Response
	VideoList []*video.Video
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "user_id get error",
		})
		return
	}
	user_id, err := strconv.ParseInt(uid.(string), 10, 64)
	// token := c.PostForm("token")
	// token校验待加
	// if _, exist := usersLoginInfo[token]; !exist {
	// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// 	return
	// }

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user_resp, err := rpc.DouyinUserInfo(context.Background(), &user.DouyinUserRequest{
		UserId: user_id,
		Token:  "test",
	})
	fmt.Println("test user", user_resp.User, user_id)
	u := user_resp.User[0]
	finalName := fmt.Sprintf("%d_%s", u.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	publich_resp, err := rpc.DouyinPublishAction(context.Background(), &video.DouyinPublishActionRequest{
		Token: "test",
		Video: &video.Video{
			Id:            2,
			Author:        u,
			PlayUrl:       saveFile,
			CoverUrl:      "./public/test.img",
			FavoriteCount: 0,
			IsFavorite:    false,
			Title:         filename,
		},
	})
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: publich_resp.BaseResp.StatusCode,
			StatusMsg:  publich_resp.BaseResp.StatusMsg,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	uid := c.Query("user_id")
	user_id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserVideoListResponse{
			Response: Response{StatusCode: errno.SuccessCode, StatusMsg: "Error parameter"},
		})
	}

	resp, err := rpc.DouyinPublishList(context.Background(), &video.DouyinPublishListRequest{
		UserId: user_id,
		Token:  "test",
	})

	c.JSON(http.StatusOK, UserVideoListResponse{
		Response: Response{
			StatusCode: errno.SuccessCode,
			StatusMsg:  resp.BaseResp.StatusMsg,
		},
		VideoList: resp.VideoList,
	})
}
