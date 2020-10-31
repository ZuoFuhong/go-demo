package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/h2non/gock.v1"
)

/**
 * gock HTTP模拟
 */
func main() {
	defer gock.Off()
	gock.New("https://www.baidu.com").Get("/").Reply(200).BodyString("MOCK DATA")

	rsp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(bytes))
}
