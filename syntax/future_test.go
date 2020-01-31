package syntax

import (
	"testing"
	"time"
)

// future 模式
// 编程中经常遇到在一个流程中需要调用多个子调用的情况，这些子调用相互之间没有依赖， 如果串行地调用，则耗时会很长，
// 此时可以使用 Go 并发编程中的 也阳re 模式 。
// futur巳 模式的基本工作原理 :
//  1)使用 chan作为函数参数。
//  2)启动 goroutine 调用函数。
//  3)通过 chan传入参数。
//  4)做其他可以并行处理的事情。
//  5)通过 chan异步获取结果。
func Test_future(t *testing.T) {
	q := query{make(chan string, 1), make(chan string, 1)}
	go execQuery(q)

	// 输入参数
	q.sql <- "select * from table"

	// 获取结果
	println(<-q.result)
}

type query struct {
	// 参数chan
	sql chan string

	// 结果chan
	result chan string
}

func execQuery(q query) {
	go func() {
		sql := <-q.sql

		// 耗时的查询逻辑
		time.Sleep(time.Second)

		q.result <- "result from " + sql
	}()
}
