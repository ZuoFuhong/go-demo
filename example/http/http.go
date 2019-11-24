package main

import (
	"fmt"
	"net/http"
)

// Http Get请求
func httpGet() {
	resp, err := http.Get("https://www.baidu.com")
	if err == nil {
		fmt.Printf("类型 = %T\n", resp)
		fmt.Printf("resp = %v", *resp)
	}
}

// Http服务器
func httpServer() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	_ = http.ListenAndServe(":8080", nil)
}

func main() {
	httpGet()
}
