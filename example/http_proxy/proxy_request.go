package http_proxy

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

// proxyRequest 使用代理
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

// pureProxyRequest TCP发HTTP协议包
func pureProxyRequest() {
	// 直出
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	// 代理
	//conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Panic(err)
	}
	// 注意：
	// 1.在http协议中使用的是 CRLF，即：\r\n
	// 2.在golang中多行文本 `` 无法理解 \r\n 转义字符。
	reqStr := "GET http://www.baidu.com/ HTTP/1.1\r\n" +
		"Host: www.baidu.com\r\n" +
		"User-Agent: curl/7.64.1\r\n" +
		"connection: close\r\n" +
		"\r\n"
	n, _ := conn.Write([]byte(reqStr))
	fmt.Printf("write n = %d\n", n)
	buf := make([]byte, 0, 4096)
	n, _ = conn.Read(buf[len(buf):cap(buf)])
	fmt.Printf("read n = %d, buf = %d\n", n, len(buf))
	fmt.Printf("%s\n", string(buf[:n]))
}
