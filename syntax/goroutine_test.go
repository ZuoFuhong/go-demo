package syntax

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

// goroutine-并发
func Test_goroutine(t *testing.T) {
	// goroutine有个特性，也就是说，如果一个goroutine没有被阻塞，那么别的goroutine就不会得到执行。
	// 这并不是真正的并发，如果你要真正的并发，需要加上下面的这行代码。

	// 设置可以同时执行的cpu的最大数量，并返回之前的设置。如果n < 1，则不改变当前设置。
	maxprocs := runtime.GOMAXPROCS(4)
	fmt.Println(maxprocs)

	// 查询本地机器上逻辑cpu的数量
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)

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
