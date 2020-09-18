package main

import (
	//"context"
	"fmt"

	"github.com/go-redis/redis/v7"
)

func main() {
	//ExampleClient()
	SentinelClient()
}

func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "cc", // no password set
		DB:       0,    // use default DB
	})
	// ctx := context.Background()
	err := client.Set("key", "hello,man", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func SentinelClient() {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"localhost:26379", "localhost:26380", "localhost:26381"},
		SentinelPassword: "ss",
		Password: "cc",
	})
	val, err := rdb.Ping().Result()
	fmt.Println(val, err)

	err = rdb.Set("key", "hello,man", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err = rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

