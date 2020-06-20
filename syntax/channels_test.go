package syntax

import (
	"fmt"
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

// 通道可以分为两个方向，一个是读，另一个是写，假如一个函数的输入参数和输出参数都是相同的 chan 类型，则该函数可以调用自己，
// 最终形成一个调用链。当然多个具有相同参数类 型的函数也能组成一个调用链，这很像 UNIX 系统的管道，是一个有类型的管道。
func Test_chan(t *testing.T) {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		// 关闭通道表示将不再在该通道上发送任何值。
		close(in)
	}()

	// 连读调用3次chain，相当于把in中的每个元素都加3
	out := chain(chain(chain(in)))
	for v := range out {
		fmt.Printf("v = %d\n", v)
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
