// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"fmt"
	"testing"
)

// 二分查找（有序数组），O(logn)的时间复杂度
func Test_BinarySearch(t *testing.T) {
	tempArr := make([]int, 100)
	for i := 0; i < len(tempArr); i++ {
		tempArr[i] = i
	}
	fmt.Println(binarySearch(tempArr, 25))
}

func binarySearch(array []int, target int) int {
	start := 0
	end := len(array) - 1
	for start <= end {
		mid := start + (end-start)/2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
