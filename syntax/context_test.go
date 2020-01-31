package syntax

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// API函数
// With 包装函数用来构建不同功能的 Context 具体对象 。
func Test_context(t *testing.T) {
	// 构造Context的room节点
	emptyContext := context.Background()
	// 1）创建一个带有退出通知的 Context 具体对象，内部创建一个 cancelCtx 的类型实例。
	context.WithCancel(emptyContext)
	// 2）创建一个带有超时通知的 Context具体对象，内部创建一个 timerCtx的类型实例。
	context.WithDeadline(emptyContext, time.Now())
	// 3）创建一个带有超时通知的 Context具体对象，内部创建一个 timerCtx 的类型实例。
	context.WithTimeout(emptyContext, 10*time.Second)
	// 4）创建一个能够传递数据的 Context具体对象，内部创建一个 valueCtx 的类型实例。
	context.WithValue(emptyContext, "name", "dazuo")
}

// 简单的案例
func Test_simple(t *testing.T) {
	ctxa, cancel := context.WithCancel(context.Background())

	// work 模拟运行并检测前端的退出通知
	go work(ctxa, "work1")

	//使用 WithDeadline 包装前面的上下文对象 ctx
	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)
	go work(ctxb, "work2")

	// 使用 WithValue 包装前面的上下文对象 ctx
	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "name", "dazuo")
	go workWithValue(ctxc, "work3")

	//故意 sleep 10 秒， 让 work2、 work3 超时退出
	time.Sleep(10 * time.Second)

	// 显式调用 workl 的 cancel 方法通知其退出
	cancel()

	//等待 work1 打印退出信息
	time.Sleep(5 * time.Second)
	fmt.Println("main stop")
}

type otherContext struct {
	context.Context
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

// 等待前端的退出通知，并试图获取 Context 传递的数据
func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("workWithValue get msg ", name)
			return
		default:
			println("value =", ctx.Value("name").(string))
			println("default ...", name)
			time.Sleep(time.Second)
		}
	}
}
