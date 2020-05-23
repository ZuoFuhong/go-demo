package syntax

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

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

// select 是类 UNIX 系统提供的 一个多路复用系统 API，Go 语言借用多路复用的概念，提供了 select关键字，用于多路监昕多个通道。
// 当监听的通道没有状态是可读或可写的， select是阻塞 的;只要监听的通道中有一个状态是可读或可写的，则 select 就不会阻塞，而是
// 进入处理就绪通道的分支流程。
// 其中：
//   1.每个 case 都必须是一个通信
//   2.所有 channel 表达式都会被求值
//   3.所有被发送的表达式都会被求值
//   4.如果任意某个通信可以进行，它就执行，其他被忽略。
//   5.如果有多个 case 都可以运行，Select 会随机公平地选取一个处理。
//   否则：
//     如果有 default 子句，则执行该语句。
//     如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
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
