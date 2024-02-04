package rediscon

import (
	"context"
	"d2-admin-service/src/infra/config"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RdbClient *redis.Client
var RdbCtx context.Context

// 初始化Redis连接池
func initRedisPool(host string, port int, password string, database int, poolSize int) *redis.Client {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port), // Redis服务器地址，如："localhost:6379"
		Password: password,                         // 如果有密码，则设置密码
		DB:       database,                         // 选择数据库（非必需，默认是0号数据库）
		PoolSize: poolSize,                         // 连接池大小
	}
	// 创建并返回一个Redis客户端实例
	client := redis.NewClient(opt)
	return client
}
func Connect() {
	applicationConfig := config.Config
	host := applicationConfig.Redis.Host
	port := applicationConfig.Redis.Port
	password := applicationConfig.Redis.Password
	database := applicationConfig.Redis.Database
	size := applicationConfig.Redis.Pool.Size

	ctx := context.Background()
	RdbClient = initRedisPool(host, port, password, database, size)
	// 检查连接是否成功
	if err := RdbClient.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	RdbCtx = ctx
	//defer RdbClient.Close()
}
