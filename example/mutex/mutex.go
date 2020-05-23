package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.Mutex(排他锁、互斥锁)
sync.Mutex一旦被锁住，其它的Lock()操作就无法再获取它的锁，只有通过Unlock()释放锁之后才能通过Lock()继续获取锁。已有的锁会
导致其它申请Lock()操作的goroutine被阻塞，且只有在Unlock()的时候才会解除阻塞。

RWMutex读写互斥锁
1.RWMutex是基于Mutex的，在Mutex的基础之上增加了读、写的信号量，并使用了类似引用计数的读锁数量
2.读锁与读锁兼容，读锁与写锁互斥，写锁与写锁互斥，只有在锁释放后才可以继续申请互斥的锁：
  - 可以同时申请多个读锁
  - 有读锁时申请写锁将阻塞，有写锁时申请读锁将阻塞
  - 只要有写锁，后续申请读锁和写锁都将阻塞
*/

var (
	m   sync.Mutex
	rwm sync.RWMutex
)

func main() {
	go applyRWLock()
	applyRWLock()

	time.Sleep(100 * time.Second)
}

func applyLock() {
	m.Lock()
	fmt.Println("hello world")
	time.Sleep(3 * time.Second)
	m.Unlock()
}

func applyRWLock() {
	rwm.RLock()
	fmt.Println("hello world")
	time.Sleep(3 * time.Second)
	rwm.RUnlock()
}
