package syntax

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

// Golang的 signal
// golang中对信号的处理主要使用os/signal包中的两个方法：一个是notify方法用来监听收到的信号；一个是 stop方法用来取消监听。

func Test_signal(t *testing.T) {
	c := make(chan os.Signal, 1)

	// 监听信号
	// 第一个参数表示接收信号的管道
	// 第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println("discovery get a signal：", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("discovery quit !!!")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
