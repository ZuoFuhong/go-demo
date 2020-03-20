package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/cms/user/login", func(w http.ResponseWriter, r *http.Request) {
		//uri:  /cms/user/login , Methods:  OPTIONS		// 浏览器跨域，预检请求
		//uri:  /cms/user/login , Methods:  POST
		fmt.Println("uri: ", r.RequestURI, ", Methods: ", r.Method)
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "*")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT,OPTIONS")
		header.Set("Content-Type", "application/json;charset=utf-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		_, _ = w.Write([]byte("{\"name\": \"dazuo\"}"))
	})
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		panic(err)
	}
}
