// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"fmt"
	"testing"
)

/*
   切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，
   或者 Python 中的 list 类型），这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，
   需要注意的是，终止索引标识的项不包括在切片内。
*/
func Test_create_Slice(t *testing.T) {
	var nameList = [3]string{"name", "age", "gender"}
	// 1.从数组或切片生成新的切片（取出元素不包含结束位置对应的值）
	var subList = nameList[1:2]
	subList[0] = "lucy"
	fmt.Printf("src = %#v, sub = %#v\n", nameList, subList)

	fmt.Println("中间至尾部：", nameList[1:])
	fmt.Println("开头至中间：", nameList[:2])
	fmt.Println("全部：", nameList[:])

	fmt.Println("********************************")

	// 2.声明新的切片
	// 除了可以从原有的数组或者切片中生成切片外，也可以声明一个新的切片，每一种类型都可以拥有其切片类型，
	// 表示多个相同类型元素的连续集合，因此切片类型也可以被声明。
	var strList []string
	fmt.Println("strList: ", strList)
	// 切片是动态结构，只能与 nil 判定相等，不能互相判定相等。
	fmt.Println(strList == nil)

	fmt.Println("********************************")

	// 3.使用 make() 函数构造切片
	// 语法：make( []Type, size, cap )
	// 其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，这个值设定后不影响 size，
	// 只是能提前分配空间，降低多次分配空间造成的性能问题。
	a := make([]int, 2, 10)
	fmt.Println("a len: ", len(a))
	fmt.Println("a value: ", a)

	fmt.Println("********************************")

	// 4.多维切片
	// 语法：var sliceName [][]...[]sliceType
	// 其中，sliceName 为切片的名字，sliceType为切片的类型，每个[ ]代表着一个维度，切片有几个维度就需要几个[ ]。
	multipSlice := [][]int{{10}, {100, 200}}
	fmt.Println(multipSlice)

	fmt.Println("********************************")

	// 5.range关键字：循环迭代切片
	// 当迭代切片时，关键字 range 会返回两个值，第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本，
	// 而不是直接返回对该元素的引用。
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	// 6.双冒号
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 新切片的内容是[3:6)，实际引用的数组是[3:8)；可以控制切片底层数组的的大小，即：cap()
	s = s[3:6:8]
	fmt.Printf("slice = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}

// 内建函数 append() 可以为切片动态添加元素
func Test_slices_append(t *testing.T) {
	// 每个切片会指向一片内存空间，这片空间能容纳一定数量的元素。当空间不能容纳足够多的元素时，切片就会进行“扩容”，返回新的切片
	ints := make([]int, 2, 2)
	ints2 := append(ints, 1)
	fmt.Printf("%#v, len = %d, cap = %d \n", ints, len(ints), cap(ints)) // 切片的长度和容量
	fmt.Printf("%#v, len = %d, cap = %d \n", ints2, len(ints2), cap(ints2))
}

// 切片复制
func Test_slices_copy(t *testing.T) {
	// 内置函数 copy() 可以将一个数组切片复制到另一个数组切片中。
	// 语法：copy( destSlice, srcSlice []T) int
	// 其中 srcSlice 为数据来源切片，destSlice 为复制的目标，目标切片必须分配过空间 且足够承载复制的元素个数，并且来源和
	// 目标的类型必须一致，copy() 函数的返回值表示实际发生复制的元素个数。
	var srcSlice = [3]int{1, 2, 3}
	var destSlice = make([]int, 2, 4)
	num := copy(destSlice[0:2], srcSlice[1:3])
	fmt.Printf("num = %d\n", num)

	srcSlice[2] = 5
	fmt.Printf("src = %#v, desc = %#v\n", srcSlice, destSlice) // output: src = [3]int{1, 2, 5}, desc = []int{2, 3}

	destSlice[1] = 6
	fmt.Printf("src = %#v, desc = %#v\n", srcSlice, destSlice) // output: src = [3]int{1, 2, 5}, desc = []int{2, 6}
}

// 将切片作为函数的参数进行传递
func Test_slices_param(t *testing.T) {
	var slice = make([]int, 2, 4)
	slice[0] = 1
	slice[1] = 2

	fmt.Println(slice)
	changeSlices(slice)
	fmt.Println(slice)
}

func changeSlices(sli []int) {
	sli[0] = 2
}
