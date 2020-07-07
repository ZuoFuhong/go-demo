package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 洗牌算法

// Collections.shuffle算法
// 该实现向后遍历列表，从最后一个元素到第二个元素，反复将随机选择的元素交换到“当前位置”。
// 元素是从从第一个元素运行到当前位置(包括当前位置)的列表中随机选择的。
func Test_CollectionShuffle(t *testing.T) {
	aList := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	for i := len(aList); i > 1; i-- {
		idx := rand.Intn(i)
		aList[i-1], aList[idx] = aList[idx], aList[i-1]
	}
	fmt.Println(aList)
}

// Fisher_Yates算法
// 1.原理
//   取两个列表，一个是洗牌前的序列A{1,2….54)，一个用来放洗牌后的序列B，B初始为空
//   随机从A取一张牌加入B末尾
// 2.复杂度
//   空间O（1），时间O（n^2)
// 3.优缺点
//   算法原理清晰，但额外开辟了一个List，而且为List删除元素是不可避免地需要移动元素
//   通过54次生成的随机数取1/54,1/53,…1/1能等概率地生成这54!种结果中的一种
func Test_FisherYates(t *testing.T) {
	const mLen = 54
	var aLen = mLen
	var aList, bList [54]int
	for a := 0; a < mLen; a++ {
		aList[a] = a + 1
	}

	rand.Seed(time.Now().UnixNano())
	for a := 0; a < mLen; a++ {
		index := a
		if aLen > 1 {
			index = rand.Intn(aLen - 1)
		}
		aLen--
		bList[a] = aList[index]
		copy(aList[index:], aList[index+1:])

		// 预览输出
		//for i := 0; i < mLen; i++ {
		//	fmt.Printf("%3d", aList[i])
		//	if i == mLen-1 {
		//		fmt.Print("\n")
		//	}
		//}
	}
	fmt.Println(bList)
}

// Knuth_Durstenfeld算法
// Knuth 和Durstenfeld 在Fisher 等人的基础上对算法进行了改进。 每次从未处理的数据中随机取出一个数字，然后把该数字放在数组的尾部，
// 即数组尾部存放的是已经处理过的数字。
// 这是一个原地打乱顺序的算法，算法时间复杂度也从Fisher算法的 O ( n 2 )提升到了 O ( n )。
func Test_KnuthDurstenfeld(t *testing.T) {
	aList := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	rand.Seed(time.Now().UnixNano())
	for a := len(aList) - 1; a > 0; a-- {
		index := rand.Intn(a)

		tmp := aList[a]
		aList[a] = aList[index]
		aList[index] = tmp
	}
	fmt.Println(aList)
}

// Inside_Out算法
// C++ stl中random_shuffle使用的就是这种算法
// 1.原理
//   在[0, i]之间随机一个下标j，然后用位置j的元素替换掉位置i的数字
//   通过54次生成的随机数取1/1, 1/2, ... 1/54能等概率地生成这54!种结果中的一种
// 2.复杂度
//   空间O（1），时间O（n)
func Test_InsideOut(t *testing.T) {
	aList := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	rand.Seed(time.Now().UnixNano())
	for a := 1; a < len(aList); a++ {
		idx := rand.Intn(a)
		aList[a], aList[idx] = aList[idx], aList[a]
	}
	fmt.Println(aList)
}
