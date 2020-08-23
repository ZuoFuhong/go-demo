package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func runHttpServer() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("request URI: %s\n", r.RequestURI) // output: /token?code=232
			fmt.Printf("request URL: %s\n", r.URL)

			bytes, _ := ioutil.ReadAll(r.Body)
			fmt.Println(string(bytes))

			_, _ = w.Write([]byte("welcome !"))
		}),
	}
	_ = server.ListenAndServe()
}

// HTTPFileServer HTTP静态资源服务器
func runHTTPFileServer() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	_ = http.ListenAndServe(":8080", nil)
}
