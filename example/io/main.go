package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 创建文件
func createFile() {
	file, e := os.Create("./temp.txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()
}

// 写文件
func writeFile() {
	file, e := os.OpenFile("./temp.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	_, e = file.WriteString("hello one\n")
	if e != nil {
		panic(e)
	}
	_, _ = file.WriteString("hello two\n")
	_, _ = file.WriteString("hello three\n")
}

// 读文件
func readFile() {
	file, e := os.Open("./temp.txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, e := reader.ReadBytes('\n')
		if e != nil && e == io.EOF {
			break
		} else if e != nil {
			panic(e)
		}
		fmt.Printf(string(line))
	}
}

func main() {
	createFile()
	writeFile()
	readFile()
}
