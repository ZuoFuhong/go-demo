package syntax

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/////////////////////////////
// Mutex互斥锁 与 RWMutex读写锁
/////////////////////////////

var mux = new(sync.Mutex)

// 互斥锁
// 1）如果对一个已经上锁的对象再次上锁，那么就会导致该锁定操作被阻塞，直到该互斥锁回到被解锁状态
// 2）可以使用defer解锁
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

// RWMutex读写锁
// 基于Mutex 实现，Lock()加写锁，Unlock()解写锁，RLock()加读锁，RUnlock()解读锁
// 多个goroutine可以同时读，读锁只会阻止写；只能一个同时写，写锁会同时阻止读写
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
