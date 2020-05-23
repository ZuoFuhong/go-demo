package syntax

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"testing"
	"time"
)

// goroutine-并发
func TestCreateGoroutineByFunc(t *testing.T) {
	// 并发执行程序
	go running()

	fmt.Println("end ...")
	time.Sleep(3 * time.Second)
}

func running() {
	defer fmt.Println("this is  defer")

	time.Sleep(1 * time.Second)
	runtime.Goexit()
	fmt.Println("hello run ...")
}

// 使用匿名函数创建goroutine
func TestCreateGoroutineByAnonymousFunc(t *testing.T) {
	go func(msg string) {
		time.Sleep(1 * time.Second)
		fmt.Println(msg)

	}("hello")

	fmt.Println("end ...")
	time.Sleep(3 * time.Second)
}

// WaitGroup-同步
var wg sync.WaitGroup
var urls = []string{
	"https://www.baidu.com",
	"https://www.baidu.com",
	"https://www.baidu.com",
}

func Test_waitGroup(t *testing.T) {
	for _, url := range urls {
		// 每一个 URL 启动一个 gorouti口e，同时给 wg 加 l
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

// 随机数生成器
// 融合了并发、缓冲、退出通知等多重特性的生成器。
func Test_generateInt(t *testing.T) {
	done := make(chan struct{})
	ch := generateInt(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 发送通知，告诉生产者停止生成
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 临时生产者已经退出
	println("NumGoroutine=", runtime.NumGoroutine())
}

func generateInt(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			// 使用select的扇入技术(Fan in)增加生成的随机源
			case ch <- <-generateIntA():
			case ch <- <-generateIntB():
			// 增加一路监听，就是对退出通知信号done的监听
			case <-done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func generateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func generateIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}
