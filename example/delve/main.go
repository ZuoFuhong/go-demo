package main

import "fmt"

/*
	Delve调试器
	目前Go语言支持GDB、LLDB和Delve几种调试器。其中GDB是最早支持的调试工具，LLDB是macOS系统推荐的标准调试工具。但是GDB和LLDB对Go语言
	的专有特性都缺乏很大支持，而只有Delve是专门为Go语言设计开发的调试工具。而且Delve本身也是采用Go语言开发。

	入门教程：https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-09-debug.html
	仓库地址：https://github.com/go-delve/delve
*/
func main() {
	nums := make([]int, 5)
	for i := 0; i < len(nums); i++ {
		nums[i] = i * i
	}
	fmt.Println(nums)
}
