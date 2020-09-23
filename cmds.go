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
	Exists(ctx context.Context, keys ...string) IntResult
	HDel(ctx context.Context, key string, fields ...string) IntResult
	HGet(ctx context.Context, key, field string) StringResult
	HGetAll(ctx context.Context, key string) StringStringMapResult
	HIncrBy(ctx context.Context, key, field string, incr int64) IntResult
	HMGet(ctx context.Context, key string, fields ...string) SliceResult
	HSet(ctx context.Context, key string, values ...interface{}) IntResult
	Incr(ctx context.Context, key string) IntResult
	Keys(ctx context.Context, pattern string) StringSliceResult
	LLen(ctx context.Context, key string) IntResult
	LPush(ctx context.Context, key string, values ...interface{}) IntResult
	LTrim(ctx context.Context, key string, start, stop int64) StatusResult
	MGet(ctx context.Context, keys ...string) SliceResult
	Ping(ctx context.Context) StatusResult
	Publish(ctx context.Context, channel string, message interface{}) IntResult
	SAdd(ctx context.Context, key string, members ...interface{}) IntResult
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) StatusResult
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) BoolResult
	SMembers(ctx context.Context, key string) StringSliceResult
	SRem(ctx context.Context, key string, members ...interface{}) IntResult
}
