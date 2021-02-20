package syntax

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/////////////////////////////
// Mutex互斥锁 与 RWMutex读写锁
/////////////////////////////

var mux = new(sync.Mutex)

func Test_CAS(t *testing.T) {
	var state int32
	swapInt := atomic.CompareAndSwapInt32(&state, 0, 1)
	fmt.Println(swapInt)
}

// sync.Mutex(排他锁、互斥锁)
// sync.Mutex一旦被锁住，其它的Lock()操作就无法再获取它的锁，只有通过Unlock()释放锁之后才能通过Lock()继续获取锁。已有的锁会
// 导致其它申请Lock()操作的goroutine被阻塞，且只有在Unlock()的时候才会解除阻塞。
func Test_Mutex(t *testing.T) {

	go printTaskName("child task")
	printTaskName("main task")

	time.Sleep(time.Second * 5)
}

func printTaskName(name string) {
	mux.Lock()
	time.Sleep(time.Second * 2)
	fmt.Println("printTaskName: " + name)
	mux.Unlock()
}

var rwmux = new(sync.RWMutex)

// RWMutex读写互斥锁
// 1.RWMutex是基于Mutex的，在Mutex的基础之上增加了读、写的信号量，并使用了类似引用计数的读锁数量
// 2.读锁与读锁兼容，读锁与写锁互斥，写锁与写锁互斥，只有在锁释放后才可以继续申请互斥的锁：
//   - 可以同时申请多个读锁
//   - 有读锁时申请写锁将阻塞
//   - 只要有写锁，后续申请读锁和写锁都将阻塞
func Test_RWMutexRead(t *testing.T) {

	go readByte("child task")
	readByte("main task")

	time.Sleep(time.Second * 4)
}

// 读操作
func readByte(name string) {
	rwmux.RLock()
	time.Sleep(time.Second * 2)
	fmt.Println("readByte name = " + name)
	rwmux.RUnlock()
}

func Test_RWMutexWrite(t *testing.T) {
	go writeByte("child task")
	writeByte("main task")

	time.Sleep(time.Second * 4)
}

// 写操作
func writeByte(name string) {
	rwmux.Lock()
	time.Sleep(time.Second * 2)
	fmt.Println("writeByte name = " + name)
	rwmux.Unlock()
}

var loadOnce sync.Once

func Test_Sync_Once(t *testing.T) {
	loadOnce.Do(func() {
		fmt.Println("once")
	})
}
