package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"time"
)

// 爬取微博搜索热榜

func main() {
	client := http.Client{Timeout: 5 * time.Second}
	req, e := http.NewRequest("GET", "https://s.weibo.com/top/summary", nil)
	if e != nil {
		panic(e)
	}

	req.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	resp, e := client.Do(req)
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()

	// html解析
	document, e := goquery.NewDocumentFromReader(resp.Body)
	if e != nil {
		panic(e)
	}
	var allData []map[string]interface{}
	document.Find(".list .list_a li").Each(func(i int, s *goquery.Selection) {
		url, exists := s.Find("a").Attr("href")
		text := s.Find("a").Find("span").Text()

		if exists {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.v2ex.com" + url})
		}
	})
	val, e := json.Marshal(allData)
	if e != nil {
		panic(e)
	}
	fmt.Printf(string(val))
}

// 输出到文件
func outputToFile(r io.Reader) {
	var buf [512]byte
	content := bytes.NewBuffer(nil)
	for {
		n, e := r.Read(buf[0:])
		content.Write(buf[0:n])
		if e != nil && e == io.EOF {
			break
		} else if e != nil {
			panic(e)
		}
	}

	file, e := os.Create("./temp.html")
	if e != nil {
		panic(e)
	}
	_, e = file.WriteString(content.String())
	if e != nil {
		panic(e)
	}
}
