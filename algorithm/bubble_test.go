// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"fmt"
	"testing"
)

// 冒泡排序
//  1.比较相邻的元素。如果第一个比第二个大，就交换他们两个。
//  2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
//  3.针对所有的元素重复以上的步骤，除了最后一个。
//  4.持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
//
// 时间复杂度：O(n^2)，最优时间复杂度：O(n)，平均时间复杂度：O(n^2)
func Test_Bubble(t *testing.T) {
	tmpArr := [10]int{2, 3, 1, 4, 6, 7, 9, 8, 5, 0}
	bubbleSort2(&tmpArr)
	fmt.Println(tmpArr)
}

func bubbleSort1(tmpArr *[10]int) {
	for i := 0; i < len(tmpArr)-1; i++ {
		for j := 0; j < len(tmpArr)-i-1; j++ {
			if tmpArr[j] > tmpArr[j+1] {
				tmp := tmpArr[j+1]
				tmpArr[j+1] = tmpArr[j]
				tmpArr[j] = tmp
			}
		}
	}
}

func bubbleSort2(tmpArr *[10]int) {
	for i := 0; i < len(tmpArr); i++ {
		for j := i + 0; j < len(tmpArr); j++ {
			if tmpArr[i] > tmpArr[j] {
				tmp := tmpArr[i]
				tmpArr[i] = tmpArr[j]
				tmpArr[j] = tmp
			}
		}
	}
}
