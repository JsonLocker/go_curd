package main

import (
	"goCurd/initializers"
	"goCurd/routers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	initializers.LoadEnv()
	initializers.LoadRedis()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()
	routers.WebRouter(Router)    // 页面路由器
	go routers.ApiRouter(Router) // api路由器
	Router.Run()
}
