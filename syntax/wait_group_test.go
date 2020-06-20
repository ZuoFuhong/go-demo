// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"net/http"
	"sync"
	"testing"
)

// WaitGroup-同步
var wg sync.WaitGroup
var urls = []string{
	"https://www.baidu.com",
	"https://www.baidu.com",
	"https://www.baidu.com",
}

func Test_waitGroup(t *testing.T) {
	for _, url := range urls {
		// 每一个 URL 启动一个 goroutine，同时给 wg 加 1
		wg.Add(1)
		go func(url string) {
			// 当前goroutine结束后给wg计数减1, wg.Done()等价于wg.Add(-1)
			defer wg.Done()

			resp, err := http.Get(url)
			if err == nil {
				println(resp.Status)
			}
		}(url)
	}
	// 等待所有请求结束
	wg.Wait()
}
