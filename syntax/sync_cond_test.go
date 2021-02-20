package syntax

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync包提供了一个条件变量类型 sync.Cond，它可以和互斥锁或读写锁（以下统称互斥锁）组合使用，用来协调
// 访问共享资源的协程。
//
// 与互斥锁不同，条件变量 sync.Cond 的主要作用并不是保证在同一时刻仅有一个线程访问某一个共享资源，而是
// 在对应的资源发生变化时，通知其它因此阻塞的线程。条件变量总是和互斥锁组合使用，互斥锁为共享资源的访问
// 提供互斥支持，而条件变量可以就共享资源的状态变化向相关线程发出通知，重在「协调」。

var sharedRsc = false

func Test_SyncCond(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait()
		}
		fmt.Println("goroutine1", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			c.Wait()
		}
		fmt.Println("goroutine2", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	// this one writes changes to sharedRsc
	time.Sleep(2 * time.Second)
	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	c.Broadcast()
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()
	wg.Wait()
}
