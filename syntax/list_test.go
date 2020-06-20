// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"container/list"
	"fmt"
	"testing"
)

/*
  	list的初始化有两种方法
	  1.通过 container/list 包的 New() 函数初始化 list
        变量名 := list.New()
	  2.通过 var 关键字声明初始化 list
        var 变量名 list.List{}
*/
func Test_ListSyntax(t *testing.T) {
	// 1.通过 container/list 包的 New() 函数初始化 list
	myList := list.New()
	// 将 fist 字符串插入到列表的尾部，此时列表是空的，插入后只有一个元素。
	myList.PushBack("first")

	// 将数值 68 放入列表。此时，列表中已经存在 fist 元素，67 这个元素将被放在 fist 的前面。
	myList.PushFront(68)

	// 保存元素句柄
	element := myList.PushFront(67)

	fmt.Println("element: ", element)
	fmt.Println(myList)

	// 获取链表首部元素 67，然后找到链表下一个元素 68
	fmt.Println(myList.Front().Next().Value)

	// 通过句柄移除元素
	myList.Remove(element)

	fmt.Println("********************************")

	// 遍历列表
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	// 2.通过 var 关键字声明初始化 list
	ageList := list.List{}
	ageList.PushBack(true)
	ageList.PushBack("dazuo")
	fmt.Print(ageList.Front().Next().Value)
}
