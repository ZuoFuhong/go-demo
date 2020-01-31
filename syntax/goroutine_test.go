package syntax

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
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

// 通道（chan）-通信
func TestChan(t *testing.T) {
	// 声明通道
	var _ chan string

	// 创建无缓冲的通道，通道存放的元素的类型为string
	ch := make(chan string)
	// 创建一个有 10 个缓冲的通道，通道存放元素的类型为 string
	// 通道分为无缓冲的通道和有缓冲的通道， Go 提供内置函数 len 和 cap，无缓冲的通道的len和cap都是 0，有缓冲的通道的
	// len 代表没有被读取的元素数，cap 代表整个通道的容量。无缓冲的通道既可以用于通信，也可以用于两个 goroutine 的同步，
	// 有缓冲的通道主要用于通信。
	//ch := make(chan string, 10)

	go func() {
		for i := 0; i < 5; i++ {
			// 发送数据
			msg := "hello i = " + strconv.Itoa(i)
			ch <- msg
			time.Sleep(time.Second)
		}
	}()

	// 循环接收数据
	for data := range ch {
		fmt.Println("接收的数据：", data)
	}

	// 关闭通道
	close(ch)
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

// select 是类 UNIX 系统提供的 一个多路复用系统 API，Go 语言借用多路复用的概念，提供 了 select关键字，用于多路监昕多个通道。
// 当监听的通道没有状态是可读或可写的， select是阻 塞 的;只要监听的通道中有 一个状态是可读或可写的，则 select 就不会阻塞，而是
// 进入处理就 绪通道的分支流程。如果监听的通道有多个可读或可写的状态， 则 select 随机选取一个处理。
func Test_select(t *testing.T) {
	ch := make(chan int, 1)
	go func(ch chan int) {
		for {
			select {
			// 0 或 1 的写入是随机的
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)

	for i := 0; i < 10; i++ {
		// 读通道c，通过通道进行同步等待
		v := <-ch
		println(v)
	}
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

// 通道可以分为两个方向，一个是读，另一个是写，假如 一个函数的输入参数和输出参数都 是相 同的 chan 类型， 则该函数可以调用自己，
// 最终形成一个调用链。当然多个具有相同参数类 型的函数也能组成一个调用链，这很像 UNIX 系统的管道，是一个有类型的管道。
func Test_chan(t *testing.T) {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	// 连读调用3次chain，相当于把in中的每个元素都加3
	out := chain(chain(chain(in)))
	for v := range out {
		println(v)
	}
}

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- 1 + v
		}
	}()
	return out
}
