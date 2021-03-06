package syntax

import (
	"sync"
	"testing"
)

/*
	sync.Pool详解
	我们通常用golang来构建高并发场景下的应用，但是由于golang内建的GC机制会影响应用的性能，为了减少GC，golang提供了对象重用的机制，
	也就是sync.Pool对象池。 sync.Pool是可伸缩的，并发安全的。其大小仅受限于内存的大小，可以被看作是一个存放可重用对象的值的容器。
	设计的目的是存放已经分配的但是暂时不用的对象，在需要用到的时候直接从pool中取。
*/

var bytePool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 10)
	},
}

func Test_syncPool(t *testing.T) {
	obj := bytePool.Get().([]byte)

	bytePool.Put(obj)
}
