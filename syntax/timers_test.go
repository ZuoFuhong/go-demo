package syntax

import (
	"fmt"
	"testing"
	"time"
)

// 定时器timer
// ticker只要定义完成，从此刻开始计时，不需要任何其他的操作，每隔固定时间都会触发（循环执行）
// timer定时器，是到固定时间后会执行一次（只会执行一次）
func Test_NewTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	go func(t *time.Timer) {
		for {
			<-t.C
			fmt.Println("get timer", time.Now().Format("2019-11-26 23:04:05"))

			// Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。如果调用时t
			// 还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			t.Reset(2 * time.Second)
		}
	}(timer)

	time.Sleep(time.Second * 30)
}

// 定时器ticker
// NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
// 整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源。
func Test_Ticker(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)
	go func(t *time.Ticker) {
		for {
			<-t.C
			fmt.Println("get ticker", time.Now().Format("2019-11-27 23:04:05"))
		}
	}(ticker)

	time.Sleep(time.Second * 30)
}
