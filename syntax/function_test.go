package syntax

import (
	"fmt"
	"strings"
	"testing"
)

/*
	函数
      1.声明函数
      2.函数变量
      3.匿名函数
      4.函数体类型实现接口

    函数值传递：在Go语言中，函数调用，参数都是以复制值传递（比较特殊的是，Go语言闭包函数对外部变量都是以引用的方式使用）
*/

// 1.声明函数
func fire(msg string) (string, int) {
	fmt.Println("hello fire")
	return msg, 1
}

func Test_Func(t *testing.T) {
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
func Test_FuncAndInterface(t *testing.T) {
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
      闭包是由函数及其相关引用环境组合而成的实体，一般通过在匿名函数中引用外部函数的局部变量。（闭包 = 函数 + 引用环境）
	  闭包对闭包外的环境引入是直接引用，编译器检测到闭包，会将闭包引用的外部变量分配到堆上。

      闭包最初的目的是减少全局变量，在函数调用的过程中隐式地传递共享变量，有其有用的一面；但是这种隐秘的共享变量的方式带来的坏处不够直接，
      不够清晰，除非是非常有价值的地方，一般不建议使用闭包。

	  函数返回的闭包引用了该函数的局部变量（参数或函数内部变量）：
	  1）多次调用该函数，返回的多个闭包所引用的外部变量是多个副本，原因是每次调用函 数都会为局部变 量分 配内存 。
      2）用一个闭包函数多次，如果该闭包修改了其引用的外部变量，则每一次调用该闭包对 该外部变量都有影响，因为闭包函数共享外部引用。

      个人理解：
      其实理解闭包的最方便的方法就是将闭包函数看成一个类，一个闭包函数调用就是实例化一个类。
	  然后就可以根据类的角度看出哪些是“全局变量”，哪些是“局部变量”了。
*/
func TestClosure(t *testing.T) {
	f := addr()
	f2 := addr()
	for i := 0; i < 3; i++ {
		fmt.Println("f：", f(2))
		fmt.Println("f2: ", f2(2))
	}
}

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/*
	函数使用示例
      1.链式处理（对数据的操作进行多步骤的处理被称为链式处理）
      2.使用匿名函数作为回调函数
*/
func Test_FuncCase(t *testing.T) {
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
