package handlers

import (
	"context"
	"strconv"

	"github.com/YuleZhang/douyin/kitex_gen/user"
	"github.com/YuleZhang/douyin/pkg/errno"

	"github.com/YuleZhang/douyin/cmd/api/rpc"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRegistResponse struct {
	Response
	UserId int64 `json:"user_id,omitempty"`
}

type UserInfoResponse struct {
	Response
	User user.User `json:"user"`
}

// Register register user info
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password, please check !"})
		return
	}
	uid, err := rpc.DouyinUserRegister(context.Background(), &user.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(http.StatusOK, UserRegistResponse{
			Response: Response{StatusCode: Err.ErrCode, StatusMsg: Err.ErrMsg},
			UserId:   uid,
		})
		return
	}
	c.JSON(http.StatusOK, UserRegistResponse{
		Response: Response{StatusCode: errno.SuccessCode, StatusMsg: "regist success!"},
		UserId:   uid,
	})
}

func UserInfo(c *gin.Context) {
	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserInfoResponse{
			Response: Response{StatusCode: errno.SuccessCode, StatusMsg: "Error parameter"},
		})
	}
	// token := c.Query("token")
	resp, err := rpc.DouyinUserInfo(context.Background(), &user.DouyinUserRequest{
		UserId: user_id,
		Token:  "test",
	})

	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: Err.ErrCode, StatusMsg: Err.ErrMsg},
		})
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		Response: Response{
			StatusCode: errno.SuccessCode,
			StatusMsg:  resp.BaseResp.StatusMsg,
		},
		User: *resp.User[0],
	})
}
