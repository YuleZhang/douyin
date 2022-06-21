package main

import (
	"net/http"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"github.com/yulezhang/douyin/cmd/api/rpc"
	"github.com/yulezhang/douyin/pkg/constants"
	"github.com/yulezhang/douyin/pkg/tracer"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()

	initRouter(r)
	// v1 := r.Group("/v1")
	// user1 := v1.Group("/douyin/user")
	// user1.POST("/login", authMiddleware.LoginHandler)
	// user1.POST("/register", handlers.Register)

	// note1 := v1.Group("/note")
	// note1.Use(authMiddleware.MiddlewareFunc())
	// note1.GET("/query", handlers.QueryNote)
	// note1.POST("", handlers.CreateNote)
	// note1.PUT("/:note_id", handlers.UpdateNote)
	// note1.DELETE("/:note_id", handlers.DeleteNote)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
