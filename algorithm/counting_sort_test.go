// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"fmt"
	"testing"
)

// 计数排序（Counting sort）是一种稳定的线性时间排序算法。计数排序不是比较排序，它的优势在于在对一定范围内的整数排序时，
// 它的复杂度为Ο(n+k)（其中k是整数的范围），快于任何比较排序算法。
//
// 1.找出待排序的数组中最大的元素。
// 2.统计数组中每个值为 i 的元素出现的次数，存入数组 C 的第 i 项。
// 3.对所有的计数累加（从 C 中的第一个元素开始，每一项和前一项相加）
// 4.反向填充目标数组：将每个元素 i 放在新数组的第 C[i] 项，每放一个元素就将 C[i] 减去1
func Test_CountSort(t *testing.T) {
	tmpArr := []int{1, 10, 2, 3, 7, 3, 4, 9, 2, 8, 6, 9, 7, 8, 10, 5, 1}
	var maxVal int
	for _, v := range tmpArr {
		if v > maxVal {
			maxVal = v
		}
	}

	countArr := make([]int, maxVal+1)
	for _, v := range tmpArr {
		countArr[v]++
	}

	newArr := make([]int, 0)
	for k, v := range countArr {
		for i := 0; v > 0 && i < v; i++ {
			newArr = append(newArr, k)
		}
	}
	fmt.Println(newArr)
}
