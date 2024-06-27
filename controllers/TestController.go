package controllers

import (
	"goCurd/helpers"
	"goCurd/initializers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 用于测试函数的方法整合体
type TestController struct{}

func (c *TestController) RedisTest(ctx *gin.Context) {
	err := initializers.Rdb.Set("api", "add cache to redis", 1*time.Hour).Err()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}
	result, err := initializers.Rdb.Get("api").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": result,
	})
}

// 测试jwt使用
func (c *TestController) JWTCheck(ctx *gin.Context) {
	str := helpers.JwtCreate("123456", 6)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": str,
	})
}

// 测试jwt解析
func (c *TestController) JwtParse(ctx *gin.Context) {
	jwt_str := helpers.JwtCreate("1234aa6", 1)
	token, err := helpers.JwtParse(jwt_str)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg":  "token解析失败",
			"data": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 500,
		"msg":  "token解析成功",
		"data": token,
	})
}
