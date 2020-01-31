package syntax

import "testing"

const (
	NUMBER = 10
)

// 固定worker工作池
// 程序中除了主要的main goroutine，还开启了如下几类goroutine
// 1）初始化任务的goroutine
// 2）分发任务的goroutine
// 3）等待所有worker结束通知，然后关闭结果通道的goroutine.

// 程序采用了三个通道，分别是：
// 1）传递task任务的通道
// 2）传递task结果的通道
// 3）接收worker处理完任务后所发送通知的通道。
type task struct {
	begin  int
	end    int
	result chan<- int
}

func Test_worker_pool(t *testing.T) {
	workers := NUMBER

	// 工作通道
	taskchan := make(chan task, 10)

	// 结果通道
	resultchan := make(chan int, 10)

	// worker信号通道
	done := make(chan struct{}, 10)

	// 初始化task的goroutine，计算100个自然数之和
	go InitTask(taskchan, resultchan, 100)

	// 分发任务到NUMBER个goroutine池
	DistributeTask(taskchan, workers, done)

	// 获取各个goroutine处理完任务的通知，并关闭通道
	go CloseResult(done, resultchan, workers)

	// 通过结果通道获取结果并汇总
	sum := ProcessResult(resultchan)

	println("sum:", sum)
}

// 任务处理：计算begin到end的和
func (t *task) do() {
	sum := 0
	for i := t.begin; i < t.end; i++ {
		sum += 1
	}
	t.result <- sum
}

func InitTask(taskchan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			b,
			e,
			r,
		}
		taskchan <- tsk
	}
	if mod != 0 {
		tsk := task{
			high + 1,
			p,
			r,
		}
		taskchan <- tsk
	}
	close(taskchan)
}

// 将task chan并分发到work goroutine处理，总的数量是workers
func DistributeTask(taskchan <-chan task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask(taskchan, done)
	}
}

// 工作goroutine处理具体工作，并将处理结果发送到结果chan
func ProcessTask(taskchan <-chan task, done chan struct{}) {
	for t := range taskchan {
		t.do()
	}
	done <- struct{}{}
}

// 通过done chan 同步等待所有工作 goroutine的结束，然后关闭结果chan
func CloseResult(done chan struct{}, resultchan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultchan)
}

// 读取结果通道，汇总结果
func ProcessResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}
