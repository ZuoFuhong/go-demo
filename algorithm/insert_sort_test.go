// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"fmt"
	"testing"
)

// 插入排序，O(n^2) 的时间复杂度，O（1）空间复杂度
//  1.从数组的第二个数据开始往前比较，即一开始用第二个数和他前面的一个比较，如果 符合条件（比前面的大或者小，自定义），则让他们交换位置。
//  2.然后再用第三个数和第二个比较，符合则交换，但是此处还得继续往前比较。
//  3.重复步骤二，一直到数据全都排完。
func Test_InsertSort(t *testing.T) {
	tmpArr := [10]int{2, 3, 1, 4, 6, 7, 9, 8, 5, 0}
	for i := 1; i < len(tmpArr); i++ {
		for j := i; j > 0; j-- {
			if tmpArr[j] < tmpArr[j-1] {
				tmpArr[j], tmpArr[j-1] = tmpArr[j-1], tmpArr[j]
			}
		}
	}
	fmt.Println(tmpArr)
}
