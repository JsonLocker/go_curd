package helpers

import (
	"goCurd/initializers"
	"os"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
)

// 生成随机字符串或者数字 type_str: byte, int, string, float, uuid, default
func Random(type_str string, length_num int) interface{} {
	var random_str = any("")
	switch type_str {
	case "byte":
		random_str = random.RandBytes(length_num)
	case "int":
		random_str = random.RandInt(1, length_num)
	case "number":
		random_str = random.RandNumeral(length_num)
	case "string":
		random_str = random.RandString(length_num)
	case "float":
		random_str = random.RandFloat(0, 1, length_num)
	case "uuid":
		random_str, _ = random.UUIdV4()
	default:
		random_str = random.RandNumeralOrLetter(length_num)
	}
	return random_str
}

// 用于返回response的json数据
func JSON(message string, data any, code int) map[string]any {
	jsonData := make(map[string]any)
	jsonData["message"] = message
	jsonData["data"] = data
	jsonData["code"] = code
	return jsonData
}

// SHA-256 加密算法
func Crypt(str string) string {
	key := os.Getenv("MMAC_SHA_KEY")
	hms := cryptor.HmacSha256(str, key)
	return hms
}

// get data from redis
func GetCache(key string) interface{} {
	result, err := initializers.Rdb.Get(key).Result()
	if err == nil {
		return result
	}
	return nil
}

// save to redis
func SetCache(redis_key string, redis_value interface{}, expire_hour int) error {
	if expire_hour == 0 {
		return initializers.Rdb.Set(redis_key, redis_value, 0).Err()
	}
	return initializers.Rdb.Set(redis_key, redis_value, time.Duration(expire_hour)*time.Hour).Err()
}
