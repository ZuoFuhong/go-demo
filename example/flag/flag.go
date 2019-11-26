package main

import (
	"flag"
	"fmt"
	"os"
)

// flag包 解析命令行参数
func main() {
	multilParse()
}

// 简单的示例
// 启动示例: $ ./main a b c d
func simpleParse() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

// 多参数解析
// 启动示例：
// $ ./flag -help
// $ ./flag -name dazuo -age 18
func multilParse() {
	var (
		name string
		age  int
	)
	flag.StringVar(&name, "name", "dazuo", "姓名")
	flag.IntVar(&age, "age", 23, "年龄")

	// 解析命令行参数
	flag.Parse()
	fmt.Println(name, age)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
}
