package atomic

import (
	"log"
	"sync"
	"testing"
)

// Go 1.3 的 sync 包中加入一个新特性：Pool
// 官方文档可以看这里 http://golang.org/pkg/sync/#Pool
// 这个类设计的目的是用来保存和复用临时对象，以减少内存分配，降低CG压力

// 还有一个重要的特性是，放进 Pool 中的对象，会在说不准什么时候被回收掉。
// 所以如果事先 Put 进去 100 个对象，下次 Get 的时候发现 Pool 是空也是有可能的。
// 不过这个特性的一个好处就在于不用担心 Pool 会一直增长，因为 Go 已经帮你在 Pool 中做了回收机制。
// 这个清理过程是在每次垃圾回收之前做的。垃圾回收是固定两分钟触发一次。
// 而且每次清理会将 Pool 中的所有对象都清理掉！
func Test_SyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return "hello empty"
		},
	}

	pool.Put("hello world")

	// Get 返回 Pool 中的任意一个对象。
	// 如果 Pool 为空，则调用 New 返回一个新创建的对象。
	log.Print(pool.Get())

	// 再取就没有了,会自动调用NEW
	log.Print(pool.Get())
}
