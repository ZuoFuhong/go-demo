package syntax

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 通道（chan）-通信
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

	// 关闭通道
	close(ch)
}

// Channel的关闭
// 关闭Channel可以通知对方内容发送完了，不用再等了
func Test_close_chan(t *testing.T) {
	channel := make(chan string)
	rand.Seed(time.Now().Unix())

	// 向channel发送随机个数的message
	go func() {
		cnt := rand.Intn(10)
		for i := 0; i < cnt; i++ {
			channel <- fmt.Sprintf("message-%2d", i)
		}
		close(channel) // 关闭Channel
	}()
	var more = true
	var msg string
	for more {
		select {
		// channel会返回两个值，一个是内容，一个是还有没有内容
		case msg, more = <-channel:
			if more {
				fmt.Println(msg)
			} else {
				fmt.Println("channel closed!")
			}
		}
	}
}
