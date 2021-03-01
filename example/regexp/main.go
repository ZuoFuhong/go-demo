package main

import (
	"fmt"
	"regexp"
)

// 正则表达式
func main() {
	buf := "abc azc a7c aac 888 a9c  tac"
	reg := regexp.MustCompile(`\S{2}`)
	ret := reg.FindString(buf)
	fmt.Printf("%s", ret)
}
