package atomic

import (
	"log"
	"sync/atomic"
	"testing"
)

// 原子类
// int32,int64,uint32,uint64,uintptr,unsafe.Pointer,共6个
// 函数的原子操作共有5种：增或减，比较并交换、载入、存储和交换它们提供了不同的功能，切使用的场景也有区别。

func Test_Atomic_int64(t *testing.T) {
	var count int64
	log.Printf("before count： %d", count)
	// 在原值的基础上 + 2
	atomic.AddInt64(&count, 2)
	log.Printf("after count：%d", count)

	// 减法
	atomic.AddInt64(&count, -1)
	log.Printf("after count2: %d", count)
}

// CAS
func Test_Atomic_CAS(t *testing.T) {
	var count int64
	// 如果addr和old相同,就用new代替addr
	swapped := atomic.CompareAndSwapInt64(&count, int64(0), int64(5))

	log.Printf("swapped: %v", swapped)
	log.Printf("after value: %d", count)
}

// Load 如果一个写操作未完成，有一个读操作就已经发生了，这样读操作使很糟糕的。
func Test_LoadInt64(t *testing.T) {
	var count int64
	// 函数atomic.LoadInt64接受一个*int64类型的指针值，并会返回该指针值指向的那个值
	// 有了“原子的”这个形容词就意味着，在这里读取value的值的同时，当前计算机中的任何CPU都不会进行其它的针对此值的读或写操作
	// &v 和 &count 内存地址不同，因此 Load 只保证读取的不是正在写入的值
	v := atomic.LoadInt64(&count)
	atomic.CompareAndSwapInt64(&v, count, count+2)
	log.Printf("after value: %d", count)
}

// 在原子的存储某个值的过程中，任何cpu都不会进行针对进行同一个值的读或写操作。
// 原子操作总会成功，因为他不必关心被操作值的旧值是什么
func Test_StoreInt64(t *testing.T) {
	var count int64
	atomic.StoreInt64(&count, 20)
	log.Printf("after value: %d", count)
}
