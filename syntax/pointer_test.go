// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"fmt"
	"testing"
)

/*
	指针语法
      1.指针地址
      2.指针类型
*/
func TestPointerSyntax(t *testing.T) {
	var hint string = "hello"

	// 对字符串取地址, ptr类型为*string
	var ptr *string = &hint

	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)

	// 打印ptr的指针地址，指针值带有0x的十六进制前缀。
	fmt.Println("ptr value：", ptr)

	// 对指针进行取值
	value := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value：%s\n", value)

	// 使用指针修改值
	*ptr = "world"
	fmt.Println("修改后的hint：", hint)

	// 使用new()创建指针
	str := new(string)
	fmt.Println("指向默认值（空字符串）：", *str)

	*str = "hello"
	fmt.Println("更新后的值：", *str)

	var a = "hello"
	var b = a
	// 输出结果：内存地址不一样
	fmt.Println("a 地址：", &a)
	fmt.Println("b 地址：", &b)
}
