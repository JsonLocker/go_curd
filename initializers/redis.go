package initializers

import (
	"os"

	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func LoadRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	Rdb = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}
