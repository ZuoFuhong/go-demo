// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"fmt"
	"testing"
	"time"
)

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

// 实现超时
func Test_timeouts(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Printf("%d\n", v)
			case v := <-time.After(2 * time.Second):
				fmt.Printf("timeout: %v\n", v)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		ch <- i
	}
}

// 使用select和默认子句来实现非阻塞的发送、接收，甚至非阻塞的多路选择
func Test_nonBlocking(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Printf("%d\n", v)
			default:
				time.Sleep(1 * time.Second)
				fmt.Printf("no message received\n")
			}
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		ch <- i
	}
}
