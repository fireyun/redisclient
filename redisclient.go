package redisclient

import (
	"context"
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error)
	Get(ctx context.Context, key string) (string, error)
}

type client struct {
	cli *redis.Client
}

func NewRedisClient(opt *redis.Options) RedisClient {
	return &client{
		cli: redis.NewClient(opt),
	}
}

func NewFailoverClient(failoverOpt *redis.FailoverOptions) RedisClient {
	return &client{
		cli: redis.NewFailoverClient(failoverOpt),
	}
}

func (c *client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return c.cli.Set(key, value, expiration).Result()
}

func (c *client) Get(ctx context.Context, key string) (string, error) {
	return c.cli.Get(key).Result()
}