package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// flag包 解析命令行参数
func main() {
	scannerStdinParse()
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
}

// 解析控制台输入
func scannerStdinParse() {
	var (
		name string
		age  int
	)
	flag.StringVar(&name, "name", "dazuo", "姓名")
	flag.IntVar(&age, "age", 23, "年龄")

	fmt.Println("等待输入...")
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		fmt.Println(line)

		//var paramList = []string{"-name", "da", "-age", "22"}

		paramList := strings.Split(line, " ")
		e := flag.CommandLine.Parse(paramList)
		if e != nil {
			fmt.Print("命令行参数解析异常！")
			return
		}
		log.Printf("name = %s age = %d\n", name, age)
	}
}
