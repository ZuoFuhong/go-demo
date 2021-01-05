package syntax

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

// 1.取消信号
//  Go语言中的 context.Context 的主要作用还是在多个 Goroutine 组成的树中同步取消信号以减少对资源的消耗和占用。
//  分析博文：https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context
func Test_CancelCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "worker")

	time.Sleep(2 * time.Second)
	// cancel内部执行关闭chan, 从而实现"取消信号"
	log.Println("closing chan")
	cancel()
	log.Println("closing done")
	time.Sleep(time.Second * 2)
}

// 任务goroutine
func worker(ctx context.Context, name string) {
	for true {
		select {
		case _, more := <-ctx.Done():
			if !more {
				log.Printf("%s chan closed\n", name)
				return
			}
		}
	}
}

// 2.使用Context传递数据
//  在真正使用传值的功能时我们也应该非常谨慎，使用 context.Context 进行传递参数请求的所有参数一种非常差的设计，
//  比较常见的使用场景是传递请求对应用户的认证令牌以及用于进行分布式追踪的请求ID。
func Test_ValueCtx(t *testing.T) {
	ctx := context.WithValue(context.Background(), "name", "dazuo")
	value := ctx.Value("name").(string)
	fmt.Println(value)
}

// 3.取消信号衍生
//   - context.WithDeadline 带有超时通知的Context，指定超时时间
//   - context.WithTimeout  带有超时通知的Context，指定超时时长
func Test_DeadlineCtx(t *testing.T) {
	ctx, cannel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cannel()

	go handle(ctx, time.Millisecond*500, "request1")
	// 第二个请求时间增加到 1500ms, 会因为上下文的过期而被终止
	go handle(ctx, time.Millisecond*1500, "request2")

	time.Sleep(time.Second * 5)
}

func handle(ctx context.Context, duration time.Duration, taskName string) {
	select {
	case <-ctx.Done():
		fmt.Printf("%s handle %v\n", taskName, ctx.Err())
	case <-time.After(duration):
		fmt.Printf("%s process request with %v \n", taskName, duration)
	}
}
