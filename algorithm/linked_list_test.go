// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

// 数组和链表
// 1.同一个数组中，所有元素的数据类型必须相同。
// 2.链表的每个元素都存储了下一个元素的地址，从而使一系列随机的内存地址串在一起。
func Test_BuiltinList(t *testing.T) {
	tempList := list.New()
	tempList.PushFront(1)
	tempList.PushBack(2)
	for v := tempList.Front(); v != nil; v = v.Next() {
		fmt.Println(v.Value)
	}

	elem := tempList.Front()
	tempList.Remove(elem)
	for v := tempList.Front(); v != nil; v = v.Next() {
		fmt.Println(v.Value)
	}
}
