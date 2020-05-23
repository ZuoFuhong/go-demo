package syntax

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/**
context With 包装函数用来构建不同功能的 Context 具体对象
*/

// 1.创建一个带有退出通知的 Context 具体对象
func Test_CancelCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "work1")

	time.Sleep(3 * time.Second)

	cancel()

	time.Sleep(time.Second * 2)
}

// 2.创建一个带有超时通知的 Context具体对象
func Test_DeadlineCtx(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	go work(ctx, "work2")

	time.Sleep(3 * time.Second)

	cancel()

	time.Sleep(time.Second * 2)
}

// 3.创建一个带有超时通知的 Context具体对象
func Test_TimeoutCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	go work(ctx, "work3")

	time.Sleep(3 * time.Second)

	cancel()

	time.Sleep(time.Second * 2)
}

// 4.创建一个能够传递数据的 Context具体对象
func Test_ValueCtx(t *testing.T) {
	ctx := context.WithValue(context.Background(), "name", "dazuo")
	value := ctx.Value("name").(string)
	fmt.Println(value)
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println("work get msg to cannel", name)
			return
		default:
			println("default ...", name)
			time.Sleep(1 * time.Second)
		}
	}
}
