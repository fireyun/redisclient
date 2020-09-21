package main

import (
	"context"
	"fmt"
	"time"

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
	rdb := redisclient.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "cc",
		DB:       0,
	})

	ctx := context.Background()
	key := "MyClient"
	listName := "mylist"
	listName2 := "mylist2"
	hashKey := "myHashKey"

	err := rdb.Set(ctx, key, "Hello,man!", 0).Err()
	checkErr(err)

	strVal, err := rdb.Get(ctx, key).Result()
	checkErr(err)
	fmt.Println("Get", key, strVal)

	interfVal, err := rdb.Eval(ctx, "return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}", []string{"key1", "key2"}, "arg1", "arg2").Result()
	checkErr(err)
	fmt.Println("Eval:", interfVal)

	statusVal, err := rdb.Ping(ctx).Result()
	checkErr(err)
	fmt.Println("Ping:", statusVal)

	rdb.Set(ctx, "key1", "value111", 0)
	rdb.Set(ctx, "key2", "value222", 0)
	interfSliVal, err := rdb.MGet(ctx, "key1", "key2").Result()
	checkErr(err)
	fmt.Println("MGet:", interfSliVal)

	intVal, err := rdb.Del(ctx, "key1", "key2").Result()
	checkErr(err)
	fmt.Println("Del:", intVal)

	sub := rdb.Subscribe(ctx, "channels")
	checkErr(err)

	go func() {
		time.Sleep(time.Second)
		rdb.Publish(ctx, "channels", "hello,a subscribe test")
	}()
	msg, err := sub.ReceiveMessage()
	checkErr(err)
	fmt.Println("ReceiveMessage:", msg)

	err = sub.Unsubscribe("channels")
	checkErr(err)

	err = sub.Close()
	checkErr(err)

	intVal, err = rdb.Incr(ctx, "key").Result()
	checkErr(err)
	fmt.Println("Incr:", "key", intVal)

	intVal, err = rdb.LPush(ctx, listName, "111").Result()
	checkErr(err)
	fmt.Println("LPush", listName, intVal)

	strSliVal, err := rdb.BRPop(ctx, time.Second*30, listName).Result()
	checkErr(err)
	fmt.Println("BRPop", listName, strSliVal)

	rdb.LPush(ctx, listName, "333")
	strVal, err = rdb.BRPopLPush(ctx, listName, listName2, time.Second).Result()
	checkErr(err)
	fmt.Println("BRPopLPush", strVal)

	intVal, err = rdb.LLen(ctx, listName2).Result()
	checkErr(err)
	fmt.Println("LLen", strVal)

	statusVal, err = rdb.LTrim(ctx, listName2, 0, 100).Result()
	checkErr(err)
	fmt.Println("LTrim", strVal)

	intVal, err = rdb.HIncrBy(ctx, hashKey, "field", 5).Result()
	checkErr(err)
	fmt.Println("HIncrBy", intVal)

	pipe := rdb.Pipeline()
	pipe.Set("aaa", 99, 0)
	pipe.Get("aaa")
	vals, err := pipe.Exec()
	checkErr(err)
	fmt.Println("Pipeline", vals)

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

	val, err := rdb.Set(ctx, key, "Hello,man!", 0).Result()
	checkErr(err)

	val, err = rdb.Get(ctx, key).Result()
	checkErr(err)
	fmt.Println(key, val)
}

func OriginClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "cc",
		DB:       0,
	})

	key := "OriginClient"
	listName := "mylist"
	listName2 := "mylist2"
	hashKey := "myHashKey"

	err := rdb.Set(key, "Hello,man!", 0).Err()
	checkErr(err)

	strVal, err := rdb.Get(key).Result()
	checkErr(err)
	fmt.Println("Get", key, strVal)

	interfVal, err := rdb.Eval("return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}", []string{"key1", "key2"}, "arg1", "arg2").Result()
	checkErr(err)
	fmt.Println("Eval:", interfVal)

	statusVal, err := rdb.Ping().Result()
	checkErr(err)
	fmt.Println("Ping:", statusVal)

	rdb.Set("key1", "value111", 0)
	rdb.Set("key2", "value222", 0)
	interfSliVal, err := rdb.MGet("key1", "key2").Result()
	checkErr(err)
	fmt.Println("MGet:", interfSliVal)

	intVal, err := rdb.Del("key1", "key2").Result()
	checkErr(err)
	fmt.Println("Del:", intVal)

	sub := rdb.Subscribe("channels")
	checkErr(err)

	go func() {
		time.Sleep(time.Second)
		rdb.Publish("channels", "hello,a subscribe test")
	}()
	msg, err := sub.ReceiveMessage()
	checkErr(err)
	fmt.Println("ReceiveMessage:", msg)

	err = sub.Unsubscribe("channels")
	checkErr(err)

	err = sub.Close()
	checkErr(err)

	intVal, err = rdb.Incr("key").Result()
	checkErr(err)
	fmt.Println("Incr:", "key", intVal)

	intVal, err = rdb.LPush(listName, "111").Result()
	checkErr(err)
	fmt.Println("LPush", listName, intVal)

	strSliVal, err := rdb.BRPop(time.Second*30, listName).Result()
	checkErr(err)
	fmt.Println("BRPop", listName, strSliVal)

	rdb.LPush(listName, "333")
	strVal, err = rdb.BRPopLPush(listName, listName2, time.Second).Result()
	checkErr(err)
	fmt.Println("BRPopLPush", strVal)

	intVal, err = rdb.LLen(listName2).Result()
	checkErr(err)
	fmt.Println("LLen", strVal)

	statusVal, err = rdb.LTrim(listName2, 0, 100).Result()
	checkErr(err)
	fmt.Println("LTrim", strVal)

	intVal, err = rdb.HIncrBy(hashKey, "field", 5).Result()
	checkErr(err)
	fmt.Println("HIncrBy", intVal)

	pipe := rdb.Pipeline()
	pipe.Set("aaa", 99, 0)
	pipe.Get("aaa")
	vals, err := pipe.Exec()
	checkErr(err)
	fmt.Println("Pipeline", vals)
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
	checkErr(err)

	val, err = rdb.Get(key).Result()
	checkErr(err)
	fmt.Println(key, val)
}

func checkErr(err error) {
	if err != nil {
		if err == redis.Nil {
			return
		}
		panic(err)
	}
}
