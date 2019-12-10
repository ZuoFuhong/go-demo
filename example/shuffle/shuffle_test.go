package shuffle

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 洗牌算法

// Fisher_Yates算法
// 取两个列表，一个是洗牌前的序列A{1,2….54)，一个用来放洗牌后的序列B，B初始为空
// 随机从A取一张牌加入B末尾
func Test_FisherYates(t *testing.T) {
	const len = 54
	var aLen = len
	var aList, bList [54]int
	for a := 0; a < len; a++ {
		aList[a] = a + 1
	}

	rand.Seed(time.Now().UnixNano())
	for a := 0; a < len; a++ {
		index := a
		if aLen > 1 {
			index = rand.Intn(aLen - 1)
		}
		aLen--
		bList[a] = aList[index]
		copy(aList[index:], aList[index+1:])

		// 预览输出
		for i := 0; i < len; i++ {
			fmt.Printf("%3d", aList[i])
			if i == len-1 {
				fmt.Print("\n")
			}
		}
	}
	fmt.Println(bList)

	// 检测
	bubbleSort(bList)
}

// Knuth_Durstenfeld算法
// Knuth 和Durstenfeld 在Fisher 等人的基础上对算法进行了改进。 每次从未处理的数据中随机取出一个数字，然后把该数字放在数组的尾部，
// 即数组尾部存放的是已经处理过的数字 。 这是一个原地打乱顺序的算法，算法时间复杂度也从Fisher算法的 O ( n 2 )提升到了 O ( n )。
func Test_KnuthDurstenfeld(t *testing.T) {
	var aList [54]int
	for a := 0; a < len(aList); a++ {
		aList[a] = a + 1
	}

	rand.Seed(time.Now().UnixNano())
	for a := len(aList) - 1; a > 0; a-- {
		index := rand.Intn(a)

		tmp := aList[a]
		aList[a] = aList[index]
		aList[index] = tmp
	}
	fmt.Println(aList)

	// 检测
	bubbleSort(aList)
}

// Inside_Out算法
// 在[0, i]之间随机一个下标j，然后用位置j的元素替换掉位置i的数字
// 通过54次生成的随机数取1/1, 1/2, ... 1/54能等概率地生成这54!种结果中的一种
// 复杂度：空间O（1），时间O（n)
func Test_InsideOut(t *testing.T) {
	var aList [54]int
	for a := 0; a < len(aList); a++ {
		aList[a] = a + 1
	}

	rand.Seed(time.Now().UnixNano())
	for a := 1; a < len(aList); a++ {
		index := rand.Intn(a)
		tmp := aList[a]
		aList[a] = aList[index]
		aList[index] = tmp
	}
	fmt.Println(aList)

	// 检测
	bubbleSort(aList)
}

// 冒泡排序
func bubbleSort(aList [54]int) {
	for i := 0; i < len(aList); i++ {
		for j := i + 1; j < len(aList); j++ {
			if aList[i] > aList[j] {
				tmp := aList[j]
				aList[j] = aList[i]
				aList[i] = tmp

			}
		}
	}
	fmt.Println(aList)
}
