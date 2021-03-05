package http_client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

var proxy = func(_ *http.Request) (*url.URL, error) {
	return url.Parse("http://127.0.0.1:9999")
}

var httpClient = http.Client{Transport: &http.Transport{Proxy: proxy}}

func proxyRequest() {
	request, err := http.NewRequest("GET", "http://www.baidu.com", strings.NewReader(""))
	if err != nil {
		log.Panic(err)
	}
	request.Header.Add("User-Agent", "curl/7.64.1")
	resp, err := httpClient.Do(request)
	if err != nil {
		log.Panic(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Status: %s\nBody: %s\n", resp.Status, string(bodyBytes))
}

func pureHttpRequest() {
	// 直出
	addr, err := net.ResolveTCPAddr("tcp", "14.215.177.39:80")
	// 代理
	//addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Panic(err)
	}
	// request
	n, _ := conn.Write([]byte(`GET / HTTP/1.1
Host: www.baidu.com
User-Agent: curl/7.64.1
Accept: */*

`))
	fmt.Printf("write n = %d\n", n)
	buf := make([]byte, 0, 4096)
	n, _ = conn.Read(buf[len(buf):cap(buf)])
	fmt.Printf("read n = %d, buf = %d\n", n, len(buf))
	fmt.Printf("%s\n", string(buf[:n]))
}
