package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
	抢红包-二倍均值算法
	剩余可用的红包金额为M，剩余人数为N，那么有如下公式：
	每次抢到的金额 = 随机区间 （0， M / N x 2）
*/
func DoubleAverage(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	// 最大可用金额
	max := amount - int64(1)*count
	// 最大可用平均值
	avg := max / count
	// 二倍均值基础上加上最小金额，防止出现 0 值
	avg2 := avg*2 + int64(1)
	// 随机红包金额序列元素，把二倍均值作为随机数的最大数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + int64(1)
	return x
}

func Test_red_packet(t *testing.T) {
	count, amount := int64(10), int64(100)
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := DoubleAverage(count-i, remain)
		remain -= x
		fmt.Println(x)
		sum += x
	}
	fmt.Print("sum: ", sum)
}
