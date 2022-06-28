package main

import (
	"net/http"

	"github.com/YuleZhang/douyin/cmd/api/rpc"
	"github.com/YuleZhang/douyin/pkg/constants"
	"github.com/YuleZhang/douyin/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

// 用viper初始化基础配置
// func initConfig() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.ReadRemoteConfig()
// 	viper.AddConfigPath("/etc/appname/")
// }

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()

	initRouter(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
