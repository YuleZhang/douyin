package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type NoteParam struct {
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func SuccessResp(c *gin.Context) {
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "success"})
}

func FailResp(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, Response{
		StatusCode: int32(code),
		StatusMsg:  err.Error()})
}
