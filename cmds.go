package redisclient

import (
	"context"
	"time"
)

// Cmds is the interface for redis commands
type Cmds interface {
	Pipeline() Pipeliner

	BRPop(ctx context.Context, timeout time.Duration, keys ...string) StringSliceResult
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) StringResult
	Del(ctx context.Context, keys ...string) IntResult
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) Result
	Get(ctx context.Context, key string) StringResult
	HIncrBy(ctx context.Context, key, field string, incr int64) IntResult
	Incr(ctx context.Context, key string) IntResult
	LLen(ctx context.Context, key string) IntResult
	LPush(ctx context.Context, key string, values ...interface{}) IntResult
	LTrim(ctx context.Context, key string, start, stop int64) StatusResult
	MGet(ctx context.Context, keys ...string) SliceResult
	Ping(ctx context.Context) StatusResult
	Publish(ctx context.Context, channel string, message interface{}) IntResult
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) StatusResult
}
