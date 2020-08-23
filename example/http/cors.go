package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/**
 * 跨域资源共享 CORS
 * Blog: https://www.cnblogs.com/maxzuo/p/13550930.html
 */
func HTTPCORS() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "x-Custome-Header,Content-Type")
		header.Set("Access-Control-Allow-Credentials", "false")
		header.Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT,OPTIONS")
		header.Set("Access-Control-Max-Age", "5")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
		}
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("RequestURI: %s, Methods: %s, ReqBody: %s\n", r.RequestURI, r.Method, reqBody)
		_, _ = w.Write([]byte("welcome !"))
	})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
