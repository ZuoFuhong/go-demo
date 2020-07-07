// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"fmt"
	"testing"
)

// 冒泡排序，O(n^2) 的时间复杂度，O（1）空间复杂度
//  1.从当前元素起，向后依次比较每一对相邻元素，若逆序则交换
//  2.对所有元素均重复以上步骤，直至最后一个元素。
func Test_Bubble(t *testing.T) {
	tmpArr := [10]int{2, 3, 1, 4, 6, 7, 9, 8, 5, 0}
	for i := 0; i < len(tmpArr)-1; i++ {
		for j := 0; j < len(tmpArr)-i-1; j++ {
			if tmpArr[j] > tmpArr[j+1] {
				tmpArr[j], tmpArr[j+1] = tmpArr[j+1], tmpArr[j]
			}
		}
	}
	fmt.Println(tmpArr)
}
