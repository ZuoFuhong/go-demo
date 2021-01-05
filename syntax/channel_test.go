package syntax

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 通道（chan）-通信
// chan   	read-write
// <-chan 	read only
// chan<- 	write only
func Test_Chan(t *testing.T) {
	// 声明通道
	var _ chan string

	// 创建无缓冲的通道，通道存放的元素的类型为string
	ch := make(chan string)

	// channel默认上是阻塞的，也就是说，如果Channel满了，就阻塞写，如果Channel空了，就阻塞读。
	//
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
}

// 关闭Channel
func Test_CloseChan(t *testing.T) {
	channel := make(chan string)
	go func() {
		rand.Seed(time.Now().Unix())
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			// 向channel发送字符串
			channel <- fmt.Sprintf("message -%2d", i)
		}
		// 关闭Channel
		close(channel)
	}()

	for true {
		select {
		// 当more为false, 表示通道已关闭
		case val, more := <-channel:
			if more {
				fmt.Println(val)
			} else {
				fmt.Println("channel closed!")
				return
			}
		}
	}
}
