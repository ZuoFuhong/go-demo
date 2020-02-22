package main

import (
	"github.com/go-redis/redis"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "47.98.199.80:6379",
		Password: "myredis", // no password set
		DB:       0,         // use default DB
	})

	val, e := client.Get("name").Result()
	if e != nil {
		panic(e)
	}
	log.Print(val)
}
