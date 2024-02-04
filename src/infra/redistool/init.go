package redistool

import (
	"d2-admin-service/src/infra/rediscon"
	"github.com/go-redis/redis/v8"
	"time"
)

func Set(key string, value string, expiration time.Duration) {
	_, err := rediscon.RdbClient.Set(rediscon.RdbCtx, key, value, expiration).Result()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	result, err := rediscon.RdbClient.Get(rediscon.RdbCtx, key).Result()
	if err == redis.Nil {
		panic("Key not found")
	} else if err != nil {
		panic(err)
	}
	return result
}

func Remove(key string) {
	_, err := rediscon.RdbClient.Del(rediscon.RdbCtx, key).Result()
	if err != nil {
		panic(err)
	}
}
