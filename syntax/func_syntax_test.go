package syntax

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

/*
	函数
      1.声明函数
      2.函数变量
      3.匿名函数
      4.函数体类型实现接口
*/

// 1.声明函数
func fire(msg string) (string, int) {
	fmt.Println("hello fire")
	return msg, 1
}

func TestFuncSyntax(t *testing.T) {
	// 2.函数变量：在 Go 语言中，函数也是一种类型，可以和其他类型一样被保存在变量中。
	f := fire
	fmt.Printf("f type: %T\n", f)
	// 调用函数
	name, age := f("dazuo")
	fmt.Printf("name = %s age = %d \n", name, age)

	fmt.Println("*******************************")

	// 3.匿名函数
	// 3.1.在定义时调用函数
	func(data int) {
		fmt.Println(data)
	}(100)

	// 3.2.将匿名函数赋值给变量
	m := func() {
		fmt.Println("this is m func")
	}
	// 调用
	m()

	// 6.defer语句（延迟调用）是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(2)
	fmt.Println(3)
}

/*
	4.接口型函数：指的是用函数实现接口，这种函数为接口型函数，这种方式适用于只有一个函数的接口。
*/
func TestFuncAndInterface(t *testing.T) {
	// 声明接口变量
	var invoker Invoker
	// 将匿名函数转为FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function:", v)
	})
	// 使用接口调用FuncCaller.Call，内部会调用函数本体
	invoker.Call("hello")
}

// 4.1.使用type关键字 定义接口
// 这个接口需要实现 Call() 方法，调用时会传入一个 interface{} 类型的变量，这种类型的变量表示任意类型的值。
type Invoker interface {
	Call(interface{})
}

// 4.2.定义一个类型（函数类型），这个类型只定义了函数的参数列表，函数参数列表与接口定义的方法一致：
type FuncCaller func(interface{})

// 4.3.然后这个类型去实现接口，实现的函数调用自己
func (f FuncCaller) Call(p interface{}) {
	// 调用f函数本体
	f(p)
}

/*
	5.闭包（Closure）---- 引用了外部变量的匿名函数
*/
func TestClosure(t *testing.T) {
	// http://c.biancheng.net/view/59.html
}

/*
	函数使用示例
      1.链式处理（对数据的操作进行多步骤的处理被称为链式处理）
      2.使用匿名函数作为回调函数
*/
func TestFuncCase(t *testing.T) {
	// 1.示例：链式处理 待处理的字符串列表
	myList := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 处理函数链
	chain := []func(string) string{
		strings.TrimSpace,
		strings.ToUpper,
	}
	stringProcess(myList, chain)

	fmt.Println("****************************")

	// 2.示例：匿名函数用作回调函数
	visit([]int{1, 2, 3}, func(i int) {
		fmt.Println(i)
	})
}

// 声明函数，使用函数作为参数
func visit(myList []int, f func(int)) {
	for _, v := range myList {
		f(v)
	}
}

// 字符串处理函数
func stringProcess(myList []string, chain []func(string) string) {
	for k, v := range myList {
		result := v

		for _, proc := range chain {
			result = proc(result)
		}

		// 将结果放回切片
		myList[k] = result
	}
}

// 7.处理运行时错误
func TestErrorSyntax(t *testing.T) {
	// 使用 errors 包进行错误的定义
	var err = errors.New("this is an error")
	// 输出错误信息
	fmt.Println(err.Error())

	// 使用自定义错误
	err = tokenError("occur error")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 自定义一个错误
func tokenError(text string) error {
	// 实例化结构体，并初始化结构体
	var errorStr = &errorString{text}
	return errorStr
}

// 定义结构体（包含错误字符串）
type errorString struct {
	s string
}

// 结构体添加方法（实现错误接口，返回错误描述）
func (e *errorString) Error() string {
	return e.s
}

// 8.服务宕机恢复
func TestPanicSyntax(t *testing.T) {

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

	// 手动触发宕机
	panic("occur panic")
}
