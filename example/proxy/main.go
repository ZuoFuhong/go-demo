package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// 反向代理
// NewSingleHostReverseProxy 返回一个新的 ReverseProxy， 将URLs 请求路由到targe的指定的scheme, host, base path 。
// 如果target的path是"/base" ，client请求的URL是 "/dir", 则target 最后转发的请求就是 /base/dir。
func main() {
	http.HandleFunc("/youtube", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("this is proxy")

		trueServer := "http://127.0.0.1:8081"
		remoteUrl, e := url.Parse(trueServer)
		if e != nil {
			_, _ = resp.Write([]byte("查询失败！"))
			return
		}
		// 转发请求
		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		proxy.ServeHTTP(resp, req)
	})
	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}

// 模拟远程服务
func mockRemoteServer() {
	http.HandleFunc("/youtube", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("this is youtute")
		_, _ = writer.Write([]byte("hello youtube"))
	})
	_ = http.ListenAndServe("127.0.0.1:8081", nil)
}
