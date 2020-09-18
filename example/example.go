package main

import (
	"context"
	"fmt"

	"github.com/fireyun/redisclient"

	"github.com/go-redis/redis/v7"
)

func main() {
	MyClient()
	MySentinelClient()
	OriginClient()
	OriginSentinelClient()
}

func MyClient() {
	rdb := redisclient.NewRedisClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "cc",
		DB:       0,
	})

	key := "MyClient"
	ctx := context.Background()

	val, err := rdb.Set(ctx, key, "Hello,man!", 0)
	if err != nil {
		panic(err)
	}

	val, err = rdb.Get(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
}

func MySentinelClient() {
	rdb := redisclient.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       "mymaster",
		SentinelAddrs:    []string{"localhost:26379", "localhost:26380", "localhost:26381"},
		SentinelPassword: "ss",
		Password:         "cc",
	})

	key := "MySentinelClient"
	ctx := context.Background()

	val, err := rdb.Set(ctx, key, "Hello,man!", 0)
	if err != nil {
		panic(err)
	}

	val, err = rdb.Get(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
}

func OriginClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "cc",
		DB:       0,
	})

	var val string
	var err error
	key := "OriginClient"

	err = rdb.Set(key, "Hello,man!", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err = rdb.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
}

func OriginSentinelClient() {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       "mymaster",
		SentinelAddrs:    []string{"localhost:26379", "localhost:26380", "localhost:26381"},
		SentinelPassword: "ss",
		Password:         "cc",
	})

	var val string
	var err error
	key := "OriginSentinelClient"

	err = rdb.Set(key, "Hello,man!", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err = rdb.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
}
