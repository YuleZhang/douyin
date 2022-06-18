package main

import (
	"context"
	"time"

	"github.com/RaymondCode/simple-demo/cmd/api/handlers"
	"github.com/RaymondCode/simple-demo/cmd/api/rpc"
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/kitex_gen/user"
	"github.com/RaymondCode/simple-demo/pkg/constants"
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
			return rpc.DouyinUserLogin(context.Background(), &user.DouyinUserLoginRequest{Username: username, Password: password})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", handlers.DouyinFeed)
	apiRouter.GET("/user/", handlers.UserInfo)
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", authMiddleware.LoginHandler)
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
