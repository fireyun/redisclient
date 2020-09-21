package redisclient

import (
	"context"
	"time"

	"github.com/go-redis/redis/v7"
)

// Client is the interface for redis client
type Client interface {
	Subscribe(ctx context.Context, channels ...string) PubSub

	Cmds
}

type client struct {
	cli *redis.Client
}

// NewClient returns a client to the Redis Server specified by Options
func NewClient(opt *redis.Options) Client {
	return &client{
		cli: redis.NewClient(opt),
	}
}

// NewFailoverClient returns a Redis client that uses Redis Sentinel for automatic failover
func NewFailoverClient(failoverOpt *redis.FailoverOptions) Client {
	return &client{
		cli: redis.NewFailoverClient(failoverOpt),
	}
}

func (c *client) Pipeline() Pipeliner {
	return c.cli.Pipeline()
}

func (c *client) BRPop(ctx context.Context, timeout time.Duration, keys ...string) StringSliceResult {
	return c.cli.BRPop(timeout, keys...)
}

func (c *client) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) StringResult {
	return c.cli.BRPopLPush(source, destination, timeout)
}

func (c *client) Del(ctx context.Context, keys ...string) IntResult {
	return c.cli.Del(keys...)
}

func (c *client) Eval(ctx context.Context, script string, keys []string, args ...interface{}) Result {
	return c.cli.Eval(script, keys, args...)
}

func (c *client) Get(ctx context.Context, key string) StringResult {
	return c.cli.Get(key)
}

func (c *client) HIncrBy(ctx context.Context, key, field string, incr int64) IntResult {
	return c.cli.HIncrBy(key, field, incr)
}

func (c *client) Incr(ctx context.Context, key string) IntResult {
	return c.cli.Incr(key)
}

func (c *client) LLen(ctx context.Context, key string) IntResult {
	return c.cli.LLen(key)
}

func (c *client) LPush(ctx context.Context, key string, values ...interface{}) IntResult {
	return c.cli.LPush(key, values...)
}

func (c *client) LTrim(ctx context.Context, key string, start, stop int64) StatusResult {
	return c.cli.LTrim(key, start, stop)
}

func (c *client) MGet(ctx context.Context, keys ...string) SliceResult {
	return c.cli.MGet(keys...)
}

func (c *client) Ping(ctx context.Context) StatusResult {
	return c.cli.Ping()
}

func (c *client) Publish(ctx context.Context, channel string, message interface{}) IntResult {
	return c.cli.Publish(channel, message)
}

func (c *client) Subscribe(ctx context.Context, channels ...string) PubSub {
	return c.cli.Subscribe(channels...)
}

func (c *client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) StatusResult {
	return c.cli.Get(key)
}
