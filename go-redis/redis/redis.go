package redis

import (
	"context"

	redisGo "github.com/redis/go-redis/v9"
)

var DB RedisPool

func InitRedis() {
	rdb := redisGo.NewClient(initOption())
	DB.defaultClient = rdb
	_, err := DB.defaultClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func initOption() *redisGo.Options {
	redisDSN := ""
	redisPASSWORD := ""
	opt, err := redisGo.ParseURL(redisDSN)
	opt.Password = redisPASSWORD
	if err != nil {
		panic(err)
	}
	return opt
}

type RedisPool struct {
	defaultClient *redisGo.Client
}
