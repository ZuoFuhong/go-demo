package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

// HTTPGet Http GET请求
func Test_HttpGet(t *testing.T) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respData))
}

// HTTPPost Post请求，格式：application/x-www-form-urlencoded
func Test_HttpPost(t *testing.T) {
	postData := url.Values{}
	postData.Add("name", "dazuo")
	postData.Add("age", "24")

	req, e := http.NewRequest("POST", "http://127.0.0.1:8080/index", strings.NewReader(postData.Encode()))
	if e != nil {
		panic(e)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{Timeout: 5 * time.Second}
	resp, e := client.Do(req)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respData))
}

// HTTPPost HTTP Post请求，格式：application/json
func Test_HttpJSONPost(t *testing.T) {
	client := http.Client{Timeout: 5 * time.Second}

	req, e := http.NewRequest("POST", "http://127.0.0.1:8080/json", bytes.NewBuffer([]byte("{\"name\": \"dazuo\"}")))
	if e != nil {
		panic(e)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, e := client.Do(req)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respData))
}
