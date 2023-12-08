package redis

import (
	"context"
	"log"
	"sync"
	"time"
)

/* 分布式锁 */
var mutex sync.Mutex

func (c *RedisPool) Lock(ctx context.Context, key string, expiration time.Duration) bool {
	mutex.Lock()
	defer mutex.Unlock()
	bool, err := c.defaultClient.SetNX(ctx, key, 1, expiration).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return bool
}

func (c *RedisPool) UnLock(ctx context.Context, key string) int64 {
	nums, err := c.defaultClient.Del(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return nums
}
