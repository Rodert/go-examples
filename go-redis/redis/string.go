package redis

import (
	"context"
	"time"

	redisGo "github.com/redis/go-redis/v9"
)

func (c *RedisPool) Get(ctx context.Context, key string) (*redisGo.StringCmd, error) {
	cmd := c.defaultClient.Get(ctx, key)
	return cmd, cmd.Err()
}

func (c *RedisPool) Set(ctx context.Context, k string, v any, exp time.Duration) error {
	cmd := c.defaultClient.Set(ctx, k, v, exp)
	return cmd.Err()
}

func (c *RedisPool) Del(ctx context.Context, k string) error {
	cmd := c.defaultClient.Del(ctx, k)
	return cmd.Err()
}

func (c *RedisPool) Exit(ctx context.Context, k string) (bool, error) {
	cmd := c.defaultClient.Exists(ctx, k)
	if cmd.Err() != nil {
		return false, cmd.Err()
	}

	if cmd.Val() == 0 {
		return false, nil
	}

	return true, nil
}

func (c *RedisPool) Expire(ctx context.Context, k string, exp time.Duration) error {
	cmd := c.defaultClient.Expire(ctx, k, exp)
	return cmd.Err()
}

func (c *RedisPool) Incr(ctx context.Context, k string) (int64, error) {
	cmd := c.defaultClient.Incr(ctx, k)
	return cmd.Val(), cmd.Err()
}
