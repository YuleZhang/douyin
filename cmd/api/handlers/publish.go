package handlers

import (
	"context"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/YuleZhang/douyin/cmd/api/rpc"
	"github.com/YuleZhang/douyin/kitex_gen/video"
	"github.com/YuleZhang/douyin/pkg/errno"
	"github.com/YuleZhang/douyin/pkg/util"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserVideoListResponse struct {
	Response
	VideoList []*video.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user_id := claims["id"]
	user_id = int64(user_id.(float64))
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if len(title) == 0 || err != nil || data == nil {
		FailResp(c, 1, errno.ParamErr)
		return
	}

	// 视频和封面存成同名，名字为uid,filename和时间戳的MD5摘要值
	relativePath := "video"
	nameBase, _ := util.MD5FileNameEncoder(user_id.(int64), title)
	fileName := nameBase + ".mp4"
	coverName := nameBase + ".jpg"
	relativeFileName := filepath.Join(relativePath, fileName)
	relativeCoverName := filepath.Join(relativePath, coverName)

	// 打开文件，转reader
	f, err := data.Open()
	defer f.Close()
	if err != nil {
		FailResp(c, 1, errno.ConvertErr(err))
		return
	}
	// 视频上传到OSS
	video_url, err := util.FileUploadToOss(relativeFileName, f)
	if err != nil {
		FailResp(c, 1, errno.ConvertErr(err))
		return
	}

	// 生成封面
	coverBytes, err := util.ReadFrameAsJpeg(video_url)
	if err != nil {
		FailResp(c, 1, errno.ConvertErr(err))
		return
	}
	// 上传封面到OSS
	cover_url, err := util.FileUploadToOss(relativeCoverName, coverBytes)
	if err != nil {
		FailResp(c, 1, errno.ConvertErr(err))
		return
	}
	// 原始数据请求到微服务
	publich_resp, err := rpc.DouyinPublishAction(context.Background(), &video.DouyinPublishActionRequest{
		UserId:   user_id.(int64),
		FileUrl:  video_url,
		CoverUrl: cover_url,
		Title:    title,
	})

	c.JSON(http.StatusOK, Response{
		StatusCode: publich_resp.BaseResp.StatusCode,
		StatusMsg:  publich_resp.BaseResp.StatusMsg,
	})
	// SuccessResp(c)
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
	})

	c.JSON(http.StatusOK, UserVideoListResponse{
		Response: Response{
			StatusCode: errno.SuccessCode,
			StatusMsg:  resp.BaseResp.StatusMsg,
		},
		VideoList: resp.VideoList,
	})
}
