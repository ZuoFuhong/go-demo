// go get github.com/gomodule/redigo/redis
package redis

import (
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	conn, e := redis.Dial("tcp", "47.98.199.80:6379")
	if e != nil {
		t.Log("Connect to redis error", e)
		return
	}
	_, _ = conn.Do("AUTH", "123456")
	s, e := redis.String(conn.Do("GET", "name"))
	t.Log(s)
}
