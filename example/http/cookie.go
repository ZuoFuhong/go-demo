package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// HTTP - 种植cookie
// Blog：https://www.cnblogs.com/maxzuo/p/13549789.html
func HttpCookie() {
	http.HandleFunc("/set_cookie", func(w http.ResponseWriter, r *http.Request) {
		c := http.Cookie{
			Name:     "user",
			Value:    "Wuhan,max",
			Domain:   "www.localhost.com",
			Path:     "/set_cookie",
			MaxAge:   10,
			Secure:   false,
			HttpOnly: true,
		}
		c2 := http.Cookie{
			Name:     "cart",
			Value:    "1,2,3",
			Domain:   "www.localhost.com",
			Path:     "/show_cookie",
			MaxAge:   10,
			Secure:   false,
			HttpOnly: true,
		}
		w.Header().Set("Set-Cookie", c.String())
		w.Header().Add("Set-Cookie", c2.String())

		_, _ = w.Write([]byte("/set_cookie：" + time.Now().String()))
	})

	http.HandleFunc("/show_cookie", func(w http.ResponseWriter, r *http.Request) {
		for _, v := range r.Cookies() {
			fmt.Printf("%s = %s\n", v.Name, v.Value)
		}

		_, _ = w.Write([]byte("/show_cookie" + time.Now().String()))
	})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
