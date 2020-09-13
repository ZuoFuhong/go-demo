package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/afex/hystrix-go/hystrix"
)

/*
	hystrix-go 是hystrix的的go语言版

	参考资料：
		仓库地址：https://github.com/afex/hystrix-go
		雪崩利器 hystrix-go 源码分析：https://www.cnblogs.com/li-peng/p/11050563.html
*/
var count int64

func init() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		MaxConcurrentRequests:  1,
		RequestVolumeThreshold: 10,
	})
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	err := hystrix.Do("my_command", func() error {
		atomic.AddInt64(&count, 1)
		fmt.Println(count)

		w.WriteHeader(200)
		_, _ = w.Write([]byte("hello world"))
		return nil
	}, func(err error) error {
		// do this when services are down
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}
