package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/juju/ratelimit"
)

/*
	1.The C10 problem
	  在同时连接到服务器的客户端数量超过 10000 个的环境中，即便硬件性能足够， 依然无法正常提供服务，简而言之，就是单机1万个并发连接问题。
	  这个概念最早由 Dan Kegel 提出并发布于其个人站点（ http://www.kegel.com/c10k.html ）

	2.wrk进行QPS测试
      $ wrk -t10 -c20 -d5s  http://127.0.0.1:8080		// 10个线程，模拟20个并发请求，持续5秒

	3.流量限制-令牌桶
	  github.com/juju/ratelimit
      原理剖析：https://chai2010.cn/advanced-go-programming-book/ch5-web/ch5-06-ratelimit.html
*/
var count int64
var bucket *ratelimit.Bucket

func init() {
	bucket = ratelimit.NewBucket(time.Minute, 10)
}

func sayHello(w http.ResponseWriter, req *http.Request) {
	if available := bucket.TakeAvailable(1); available > 0 {
		atomic.AddInt64(&count, 1)
		fmt.Println(count)
	}

	w.WriteHeader(200)
	_, _ = w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}
