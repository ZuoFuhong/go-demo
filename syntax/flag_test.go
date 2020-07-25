package syntax

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

// 获取启动参数: $ ./main a b c d
func osArgs() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

// 命令行参数解析
// $ ./flag -help
// $ ./flag -name dazuo -age 18
func Test_Flag(t *testing.T) {
	var (
		name string
		age  int
	)
	flag.StringVar(&name, "name", "", "姓名")
	flag.IntVar(&age, "age", 0, "年龄")

	paramList := []string{"-name", "dazuo", "-age", "22"}
	err := flag.CommandLine.Parse(paramList)
	if err != nil {
		panic(err)
	}
	fmt.Printf("name = %s, age = %d", name, age)
}
