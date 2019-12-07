package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

// Http Get请求
func Test_HttpGet(t *testing.T) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Print(result.String())
}

// Http Post请求，格式：application/json
func Test_HttpJSONPost(t *testing.T) {
	client := http.Client{Timeout: 5 * time.Second}

	req, e := http.NewRequest("POST", "http://127.0.0.1:8080", bytes.NewBuffer([]byte("{\"name\": \"dazuo\"}")))
	if e != nil {
		panic(e)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, e := client.Do(req)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()
	var buf [512]byte
	content := bytes.NewBuffer(nil)
	for {
		n, e := resp.Body.Read(buf[0:])
		content.Write(buf[0:n])
		if e != nil && e == io.EOF {
			break
		} else if e != nil {
			panic(e)
		}
	}
	fmt.Printf(content.String())
}

// Http Post请求，格式：application/x-www-form-urlencoded
func Test_HttpPost(t *testing.T) {
	postData := url.Values{}
	postData.Add("cmd", "base/tomcatPressureTest")
	postData.Add("data", "{}")

	req, e := http.NewRequest("POST", "http://192.168.3.185/zxcity_restful/ws/rest", strings.NewReader(postData.Encode()))
	if e != nil {
		panic(e)
	}
	req.Header.Set("User-Agent", "go/1.12.4")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apikey", "test")

	client := http.Client{Timeout: 5 * time.Second}
	resp, e := client.Do(req)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()

	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Printf(result.String())
}
