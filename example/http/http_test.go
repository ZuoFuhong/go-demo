package example

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"
)

// Http Get请求
func Test_HttpGet(t *testing.T) {
	resp, err := http.Get("https://www.baidu.com")
	if err == nil {
		fmt.Printf("类型 = %T\n", resp)
		fmt.Printf("resp = %v", *resp)
	}
}

// Http文件服务器
func Test_HttpFileServer(t *testing.T) {
	http.Handle("/", http.FileServer(http.Dir(".")))
	_ = http.ListenAndServe(":8080", nil)
}

// http服务器
func Test_HttpServer(t *testing.T) {
	http.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {
		// 设置上下文参数（WithValue返回父节点的一个副本）
		cx := context.WithValue(req.Context(), "param", 88)
		// 提取上下文参数
		if val, ok := cx.Value("param").(int); ok {
			fmt.Printf("param: %d\n", val)
		}

		fmt.Println("hello")
		headers := req.Header
		fmt.Println(headers["Cookie"])

		_, _ = resp.Write([]byte("hello client"))
	})
	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}

// http服务器
func Test_HttpServer2(t *testing.T) {
	server := http.Server{
		Addr: "127.0.0.1:3031",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("hello client"))
		}),
	}
	_ = server.ListenAndServe()
}

// 反向代理
// NewSingleHostReverseProxy 返回一个新的 ReverseProxy， 将URLs 请求路由到targe的指定的scheme, host, base path 。
// 如果target的path是"/base" ，client请求的URL是 "/dir", 则target 最后转发的请求就是 /base/dir。
func Test_MockProxyServer(t *testing.T) {
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
func Test_MockRemoteServer(t *testing.T) {
	http.HandleFunc("/youtube", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("this is youtute")
		_, _ = writer.Write([]byte("hello youtube"))
	})
	_ = http.ListenAndServe("127.0.0.1:8081", nil)
}
