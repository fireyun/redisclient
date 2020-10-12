package redisclient

import (
	"context"
	"time"
)

// Commands is the interface for redis commands
type Commands interface {
	Pipeline() Pipeliner

	BRPop(ctx context.Context, timeout time.Duration, keys ...string) StringSliceResult
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) StringResult
	Close() error
	Del(ctx context.Context, keys ...string) IntResult
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) Result
	Expire(ctx context.Context, key string, expiration time.Duration) BoolResult
	FlushDB(ctx context.Context) StatusResult
	Get(ctx context.Context, key string) StringResult
	Exists(ctx context.Context, keys ...string) IntResult
	HDel(ctx context.Context, key string, fields ...string) IntResult
	HGet(ctx context.Context, key, field string) StringResult
	HGetAll(ctx context.Context, key string) StringStringMapResult
	HIncrBy(ctx context.Context, key, field string, incr int64) IntResult
	HKeys(ctx context.Context, key string) StringSliceResult
	HMGet(ctx context.Context, key string, fields ...string) SliceResult
	HSet(ctx context.Context, key string, values ...interface{}) IntResult
	Incr(ctx context.Context, key string) IntResult
	Keys(ctx context.Context, pattern string) StringSliceResult
	LLen(ctx context.Context, key string) IntResult
	LPush(ctx context.Context, key string, values ...interface{}) IntResult
	LRange(ctx context.Context, key string, start, stop int64) StringSliceResult
	LTrim(ctx context.Context, key string, start, stop int64) StatusResult
	MGet(ctx context.Context, keys ...string) SliceResult
	MSet(ctx context.Context, values ...interface{}) StatusResult
	Ping(ctx context.Context) StatusResult
	Publish(ctx context.Context, channel string, message interface{}) IntResult
	RPop(ctx context.Context, key string) StringResult
	RPopLPush(ctx context.Context, source, destination string) StringResult
	RPush(ctx context.Context, key string, values ...interface{}) IntResult
	SAdd(ctx context.Context, key string, members ...interface{}) IntResult
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) StatusResult
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) BoolResult
	SMembers(ctx context.Context, key string) StringSliceResult
	SRem(ctx context.Context, key string, members ...interface{}) IntResult
	TTL(ctx context.Context, key string) DurationResult
}
