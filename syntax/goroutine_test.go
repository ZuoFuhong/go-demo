package syntax

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 使用普通函数创建 goroutine
func TestCreateGoroutineByFunc(t *testing.T) {
	// 并发执行程序
	go running()

	fmt.Println("end ...")
	time.Sleep(3 * time.Second)
}

func running() {
	time.Sleep(1 * time.Second)
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

// 通道（chan）
func TestChan(t *testing.T) {
	// 声明通道
	var _ chan string

	// 创建通道
	ch := make(chan string)

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
