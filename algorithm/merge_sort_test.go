// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"fmt"
	"testing"
)

// 归并排序，O(nlogn)的时间复杂度，O(n)的空间复杂度
// 归并排序是用分治法（Divide and Conquer），分治模式在每一层递归上有三个步骤：
//  1.分解（Divide）：将n个元素分成个含n/2个元素的子序列。
//  2.解决（Conquer）：用合并排序法对两个子序列递归的排序。
//  3.合并（Combine）：合并两个已排序的子序列已得到排序结果。
func Test_MergeSort(t *testing.T) {
	tempArr := []int{2, 3, 1, 4, 6, 7, 9, 8, 5, 0}
	mergeSort(tempArr, 0, len(tempArr)-1)
	fmt.Println(tempArr)
}

func mergeSort(array []int, start, end int) {
	if start < end {
		mid := (start + end) / 2
		mergeSort(array, start, mid)
		mergeSort(array, mid+1, end)

		merge(array, start, mid, end)
	}
}

func merge(array []int, start, mid, end int) {
	tempArr := make([]int, end-start+1)
	p := 0
	p1 := start
	p2 := mid + 1

	for p1 <= mid && p2 <= end {
		if array[p1] <= array[p2] {
			tempArr[p] = array[p1]
			p++
			p1++
		} else {
			tempArr[p] = array[p2]
			p++
			p2++
		}
	}
	for p1 <= mid {
		tempArr[p] = array[p1]
		p++
		p1++
	}
	for p2 <= end {
		tempArr[p] = array[p2]
		p++
		p2++
	}
	for i := 0; i < len(tempArr); i++ {
		array[i+start] = tempArr[i]
	}
}
