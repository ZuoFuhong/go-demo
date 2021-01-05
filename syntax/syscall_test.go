// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"testing"
)

// 系统调用主要通过两个包完成的。一个是os包，一个是syscall包。

func Test_ENV(t *testing.T) {
	_ = os.Setenv("FOO", "1")
	fmt.Println("FOO: ", os.Getenv("FOO"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}

// 通过Go程序生成一个pygmentize进程（生成外部进程）
func Test_spawning_processes(t *testing.T) {
	dateCmd := exec.Command("date")
	output, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

// exec会执行参数指定的命令，但是并不创建新的进程，只在当前进程空间内执行，
// 即替换当前进程的执行内容，他们重用同一个进程号PID。
func Test_exec(t *testing.T) {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// Golang的 signal
// golang中对信号的处理主要使用os/signal包中的两个方法：一个是notify方法用来监听收到的信号；一个是 stop方法用来取消监听。
func Test_Signal(t *testing.T) {
	c := make(chan os.Signal, 1)

	// 监听信号
	// 第一个参数表示接收信号的管道
	// 第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println("get a signal：", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("quit !!!")
			return
		case syscall.SIGHUP:
			// 当session关闭时, 忽略SIGHUP信号, 保持守护进程继续运行
		default:
			return
		}
	}
}
