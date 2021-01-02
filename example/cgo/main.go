package main

//#cgo CFLAGS: -I./number
//#cgo LDFLAGS: -L${SRCDIR}/number -lnumber
//
//#include "number.h"
import "C"
import "fmt"

// CGO使用静态库
// $ cd ./number
// $ gcc -c -o number.o number.c
// $ ar rcs libnumber.a number.o
//
// 参考：https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-09-static-shared-lib.html
func main() {
	fmt.Println(C.number_add_mod(10, 5, 12))
}
