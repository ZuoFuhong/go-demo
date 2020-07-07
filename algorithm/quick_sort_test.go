// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"fmt"
	"testing"
)

// 快速排序，O(nlogn) 的时间复杂度，在分治过程中，以基准元素为中心，把其它元素移动到它的两边。
//  - 基准元素的选择：随机选择一个元素作为基准元素，并且让基准元素和数列首元素交换位置
//  - 单边循环法、双边循环法；基于递归和基于非递归（绝大数递归的逻辑，可以用栈来实现）
func Test_QuickSort(t *testing.T) {
	tmpArr := []int{2, 3, 1, 4, 6, 7, 9, 8, 5, 0}
	quickSort(tmpArr, 0, len(tmpArr)-1)
	fmt.Println(tmpArr)
}

func quickSort(tmpArr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	// 得到基准元素的位置
	pivotIndex := partition(tmpArr, startIndex, endIndex)
	quickSort(tmpArr, startIndex, pivotIndex-1)
	quickSort(tmpArr, pivotIndex+1, endIndex)
}

func partition(tmpArr []int, startIndex, endIndex int) int {
	pivot := tmpArr[startIndex]
	mark := startIndex

	for i := startIndex + 1; i <= endIndex; i++ {
		// 如果遍历的元素小于基准元素，则首先将 mark 指针右移1位，然后将 最新遍历到的元素
		// 和 mark 指针所在位置的元素交换位置。
		if tmpArr[i] < pivot {
			mark++
			tmpArr[mark], tmpArr[i] = tmpArr[i], tmpArr[mark]
		}
	}
	tmpArr[startIndex] = tmpArr[mark]
	tmpArr[mark] = pivot
	return mark
}
