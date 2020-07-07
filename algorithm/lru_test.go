// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

// 缓存淘汰算法--LRU算法
// LRU（Least recently used，最近最少使用）算法根据数据的历史访问记录来进行淘汰数据。
//  1.新数据插入到链表头部；
//  2.每当缓存命中（即缓存数据被访问），则将数据移到链表头部；
//  3.当链表满的时候，将链表尾部的数据丢弃。

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type Pair struct {
	key   int
	value int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c LRUCache) Get(key int) int {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (c LRUCache) Put(key, value int) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value = Pair{key: key, value: value}
	} else {
		if c.list.Len() >= c.capacity {
			delete(c.cache, c.list.Back().Value.(Pair).key)
			c.list.Remove(c.list.Back())
		}
		c.list.PushFront(Pair{key: key, value: value})
		c.cache[key] = c.list.Front()
	}
}

func Test_LRU(t *testing.T) {
	lruCache := NewLRUCache(10)
	lruCache.Put(1, 10)
	lruCache.Put(2, 20)
	fmt.Println(lruCache.Get(1))
}
