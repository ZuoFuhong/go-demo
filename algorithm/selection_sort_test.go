// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"fmt"
	"testing"
)

// 选择排序，O(n^2)时间复杂度，O（1）空间复杂度
// 首先在未排序序列中找到最小元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小元素，
// 然后放到已排序序列的末尾。
func Test_SelectionSort(t *testing.T) {
	tmpArr := [10]int{2, 3, 1, 4, 6, 9, 8, 7, 5, 0}
	for i := 0; i < len(tmpArr)-1; i++ {
		min := i
		for j := i + 1; j < len(tmpArr); j++ {
			if tmpArr[j] < tmpArr[min] {
				min = j
			}
		}
		tmpArr[i], tmpArr[min] = tmpArr[min], tmpArr[i]
	}
	fmt.Println(tmpArr)
}
