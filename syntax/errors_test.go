// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

// 1.处理运行时错误
func TestErrorSyntax(t *testing.T) {
	// 使用 errors 包进行错误的定义
	var err = errors.New("this is an error")
	// 输出错误信息
	fmt.Println(err.Error())
}

// 2.自定义一个错误（实现错误接口，返回错误描述）
type TokenError struct {
	ErrMsg string
}

func (e TokenError) Error() string {
	return e.ErrMsg
}

func Test_TokenError(t *testing.T) {
	var err error = TokenError{"Token is invalid."}
	switch err.(type) {
	case TokenError:
		fmt.Println("1")
	case runtime.Error:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}
}

// 3.服务宕机恢复
func Test_Panic(t *testing.T) {

	// defer后边会接一个函数，但该函数不会立刻被执行，而是等到包含它的程序返回时(包含它的函数执行了return语句、运行到函数结尾
	// 自动返回、对应的goroutine发生panic）defer函数才会被执行。通常用于资源释放、打印日志、异常捕获等。
	defer func() {
		// Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效，
		// 在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，如果当前的 goroutine 陷入恐慌，调用 recover
		// 可以捕获到 panic 的输入值，并且恢复正常的执行。
		err := recover()

		// 使用type关键字 做类型开关
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error", err)
		default:
			// 非运行时错误
			fmt.Println("error: ", err)
		}
	}()

	// panic内置函数停止当前goroutine的正常执行，当函数F调用panic时，函数F的正常执行被立即停止， 然后运行所有在F函数中的defer
	// 函数，然后F返回到调用他的函数对于调用者G，F函数的行为就像panic 一样，终止G的执行并运行G中所defer函数，此过程会一直继续执行
	// 到goroutine所有的函数。panic可以通过内置的recover来捕获。
	panic("occur panic")
}
