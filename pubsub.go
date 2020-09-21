package redisclient

import (
	"time"

	"github.com/go-redis/redis/v7"
)

// PubSub is the interface for redis Pub/Sub commands
type PubSub interface {
	Channel() <-chan *redis.Message
	ChannelSize(size int) <-chan *redis.Message
	ChannelWithSubscriptions(size int) <-chan interface{}
	Close() error
	PSubscribe(patterns ...string) error
	PUnsubscribe(patterns ...string) error
	Ping(payload ...string) error
	Receive() (interface{}, error)
	ReceiveMessage() (*redis.Message, error)
	ReceiveTimeout(timeout time.Duration) (interface{}, error)
	String() string
	Subscribe(channels ...string) error
	Unsubscribe(channels ...string) error
}
