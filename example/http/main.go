package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

func main() {
	simpleHttpServer()
}

func simpleHttpServer() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("requestUri: %s\n", r.RequestURI)
			fmt.Printf("headers: %s\n", r.Header)

			var buf [512]byte
			content := bytes.NewBuffer(nil)
			for {
				n, err := r.Body.Read(buf[0:])
				content.Write(buf[0:n])
				if err != nil && err == io.EOF {
					break
				} else if err != nil {
					panic(err)
				}
			}
			fmt.Printf("content: %s\n", content.String())

			_, _ = w.Write([]byte("hello client"))
		}),
	}
	_ = server.ListenAndServe()
}

// Http服务器
func httpServer() {
	// 静态资源
	http.Handle("/", http.FileServer(http.Dir(".")))
	// 注册处理函数
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
	_ = http.ListenAndServe(":8080", nil)
}
