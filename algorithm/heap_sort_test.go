// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"fmt"
	"testing"
)

/*
	二叉堆
     二叉堆本质上是一种完全二叉树，它分为两个类型：
	 1.最大堆：最大堆任何一个父节点的值，都大于等于它左右孩子节点的值。
	 2.最小堆：最小堆任何一个父节点的值，都小于等于它左右孩子节点的值。

	堆的自我调整
	 1.插入节点：二叉堆的节点插入，插入位置是完全二叉树的最后一个位置，然后和它的父节点比较，如果 小于/大于 父节点，
       则让新节点"上浮"，和父节点交换位置。
	 2.删除节点：二叉堆的节点删除过程和插入过程正好相反，所删除的是处于堆顶的节点。这时候，为了维持完全二叉树的结构，
       我们把堆的最后一个节点10补到原本堆顶的位置。然后进行下沉调整。
	 3.构建二叉堆：将一个无序的完全二叉树调整为二叉堆，本质上就是让所有非叶子节点依次下沉。

	注意：
	 二叉堆是一个完全二叉树，存储方式是顺序存储，存储在数组中，完全二叉树存储在数组中，不会有空位置。可以根据数组的
     下标来计算。假设父节点的下标是parent，那么它的左孩子的下标就是2 * parent + 1，右孩子就是2 * parent + 2。

	漫画：什么是二叉堆？：https://mp.weixin.qq.com/s?__biz=MzI1MTIzMzI2MA==&mid=2650562926&idx=1&sn=4ad824f145b890ff68f35c34f1d9a164&chksm=f1fed7edc6895efb200151d82a4180a925ae28da1f9fe586b7c82526d6f074468181a9babc0e&scene=21#wechat_redirect
*/

/*
	堆排序的步骤：
	 1.将无序数组构建成最大二叉堆。
	 2.循环删除栈顶元素，并将该元素移到集合尾部，调整堆产生新的堆顶。
	   第一步，把无序数组构建成 最大二叉堆，这一步的时间复杂度是 O(n)。
	   第二步，需要进行 n - 1 次循环。每次循环调用一次 downAdjust 方法，所以第二步的计算规模是 (n-1) * logn，
       时间复杂度为 O(nlogn)。
*/
func Test_HeapSort(t *testing.T) {
	tmpArr := []int{2, 3, 1, 4, 6, 7, 9, 8, 5, 0}
	buildHeap(tmpArr)
	fmt.Println(tmpArr)

	heapSort(tmpArr)
	fmt.Println(tmpArr)
}

// 堆排序（升序）
func heapSort(tmpArr []int) {
	// 循环删除栈顶元素，然后下沉调整
	for i := len(tmpArr) - 1; i > 0; i-- {
		tmpArr[i], tmpArr[0] = tmpArr[0], tmpArr[i]
		downAdjust(tmpArr, 0, i)
	}
}

// 构建堆
func buildHeap(tmpArr []int) {
	// 从最后一个非叶子节点开始，依次下沉调整
	// parentIndex = (len(tmpArr) - 2) / 2 <==> 2 * parentIndex + 2 = len(tmpArr)
	for i := (len(tmpArr) - 2) / 2; i >= 0; i-- {
		downAdjust(tmpArr, i, len(tmpArr))
	}
}

// "下沉"调整最大堆
func downAdjust(tmpArr []int, parentIndex int, length int) {
	childIndex := 2*parentIndex + 1
	for childIndex < length {
		// 如果有右孩子，且右孩子小于左孩子的值，则定位到右孩子
		if childIndex+1 < length && tmpArr[childIndex+1] > tmpArr[childIndex] {
			childIndex++
		}
		if tmpArr[parentIndex] >= tmpArr[childIndex] {
			break
		}
		tmpArr[parentIndex], tmpArr[childIndex] = tmpArr[childIndex], tmpArr[parentIndex]
		// 继续下沉
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
}

// 首先构建堆，然后插入节点，依次上浮
func Test_Insert(t *testing.T) {
	tmpArr := []int{2, 3, 1, 4, 6, 7, 9, 8, 5, 0}
	buildHeap(tmpArr)

	// 插入节点
	tmpArr = append(tmpArr, 10)
	upAdjust(tmpArr)
	fmt.Println(tmpArr)
}

// "上浮"调整最大堆
func upAdjust(tmpArr []int) {
	childIndex := len(tmpArr) - 1
	parentIndex := (childIndex - 1) / 2
	for childIndex > 0 && tmpArr[childIndex] > tmpArr[parentIndex] {
		tmpArr[childIndex], tmpArr[parentIndex] = tmpArr[parentIndex], tmpArr[childIndex]
		childIndex = parentIndex
		parentIndex = (parentIndex - 1) / 2
	}
}
