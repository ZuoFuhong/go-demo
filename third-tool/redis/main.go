package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "myredis",
		DB:       0,
	})
	err := client.Set("name", "dazuo", time.Second*10).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
