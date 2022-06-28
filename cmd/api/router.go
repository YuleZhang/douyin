package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/YuleZhang/douyin/cmd/api/handlers"
	"github.com/YuleZhang/douyin/cmd/api/rpc"
	"github.com/YuleZhang/douyin/controller"
	"github.com/YuleZhang/douyin/kitex_gen/user"
	"github.com/YuleZhang/douyin/pkg/constants"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			username := c.Query("username")
			password := c.Query("password")
			if len(username) == 0 || len(password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			uid, err := rpc.DouyinUserLogin(context.Background(), &user.DouyinUserLoginRequest{Username: username, Password: password})
			c.Set("user_id", uid)
			return uid, err
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt, form: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			uid, exists := c.Get("user_id")
			if !exists {
				c.JSON(http.StatusOK, gin.H{
					"status_code": http.StatusBadRequest,
					"status_msg":  "login failed",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "login success",
				"user_id":     uid,
				"token":       token,
				"expire":      expire.Format(time.RFC3339),
			})
		},
	})

	// public directory is used to serve static resources, take care that the path is current directory
	r.Static("/static", "../../public")

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claim := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claim)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// token is not necessary in those api
	r.GET("/douyin/feed/", handlers.DouyinFeed)
	r.POST("/douyin/user/register/", handlers.Register)
	r.POST("/douyin/user/login/", authMiddleware.LoginHandler)

	// token are required for those api
	apiRouter := r.Group("/douyin")
	apiRouter.Use(authMiddleware.MiddlewareFunc())

	// basic apis
	apiRouter.GET("/user/", handlers.UserInfo)
	apiRouter.POST("/publish/action/", handlers.Publish)
	apiRouter.GET("/publish/list/", handlers.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
