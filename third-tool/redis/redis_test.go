package redis

import (
	"github.com/go-redis/redis"
	"log"
	"testing"
)

func Test_Redis(t *testing.T) {
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
