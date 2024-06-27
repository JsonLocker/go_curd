package routers

import (
	"goCurd/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRouter(router *gin.Engine) {

	// api
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "api"})
	})

	testRouter := router.Group("/api/test/")
	{
		test := controllers.TestController{}
		testRouter.GET("/jwt", test.JWTCheck)
		testRouter.GET("/jwt2", test.JwtParse)
		testRouter.GET("/redis", test.RedisTest)
	}

}
