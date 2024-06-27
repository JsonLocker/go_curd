package routers

import (
	"goCurd/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func WebRouter(router *gin.Engine) {
	router.Static("/asset", "./assets") //前路由 后物理地址
	router.LoadHTMLGlob("views/**/*")

	// 首页
	router.GET("/", func(c *gin.Context) {
		port := os.Getenv("DB_HOST")
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "Welcome to Golang Gin Web Framework " + port,
		})
	})

	// 文章路由
	posts := router.Group("/post")
	{
		post := controllers.PostController{}
		posts.GET("/", post.Index)
		posts.GET("/create", post.Create)
		posts.GET("/:id", post.Show)
		posts.GET("/edit/:id", post.Edit)
		posts.POST("/store/:id", post.Store)
		posts.POST("/:id", post.Delete)
	}

	// websocket 路由
	socketRouter := router.Group("/ws")
	{
		ws := controllers.WebSocketController{}
		socketRouter.GET("/", ws.Index)
		socketRouter.GET("/chat", ws.Chat)
		socketRouter.GET("/mqtt", ws.Mqtt)
	}

}
