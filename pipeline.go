package redisclient

import (
	"github.com/go-redis/redis/v7"
)

// Pipeliner is interface for redis pipeline technique
type Pipeliner interface {
	redis.Cmdable
	Close() error
	Discard() error
	Exec() ([]redis.Cmder, error)
}
