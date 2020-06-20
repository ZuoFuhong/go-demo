// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"fmt"
	"testing"
)

/*
	数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是固定的。
*/
func Test_Array(t *testing.T) {
	// var 数组变量名 [元素数量]T
	var arr [4]int

	// 数组长度
	fmt.Println("arr length: ", len(arr))
	// 默认值
	fmt.Println("arr value：", arr)

	// 数组赋值，数组越界，运行时会报出宕机
	arr[0] = 1
	fmt.Println("arr value: ", arr)

	// 需要确保大括号后面的元素数量与数组的大小一致
	var nameList = [3]string{"dazuo", "wang", "li"}
	fmt.Println(nameList)

	// 让编译器在编译时，根据元素个数确定数组大小
	var ageList = [...]int{1, 3, 4}
	fmt.Println(ageList)
	fmt.Println("ageList length: ", len(ageList))

	// 比较两个数组是否相等
	// 如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（== 和!=）来判断两个数组是否相等，
	// 只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组，否则程序将无法完成编译。
	aList := [...]int{1, 2, 3}
	bList := [3]int{1, 2, 3}
	fmt.Println("判断数据是否相等：", aList == bList) // output: true

	// 遍历数组
	for k, v := range aList {
		fmt.Println(k, v)
	}
	// 多维数组
	var multiArr = [2][2]string{{"dazuo", "age"}, {"jin", "age"}}
	fmt.Println(multiArr)

	// 1.Go中的数组是值类型，换句话说，如果你将一个数组赋值给另外一个数组，那么，实际上就是将整个数组拷贝一份
	// 2.如果Go中的数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针。因此，在Go中如果
	//   将数组作为函数的参数传递的话，那效率就肯定没有传递指针高了。
	cList := aList
	cList[0] = 3
	fmt.Println("aList: ", aList) // output: [1, 2, 3]

	changeArr(aList)
	fmt.Println("aList: ", aList) // output: [1, 2, 3]
}

func changeArr(arr [3]int) {
	arr[0] = 4
}
