// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

// 二叉查找树
// 1.每个结点的键都大于其左子树中的任意结点的键而小于右子树的任意结点的键。
// 2.任意结点的左、右子树也分别为二叉搜索树。
func Test_BinarySearchTree(t *testing.T) {
	bst := new(BST)
	bst.Put(5, "5")
	bst.Put(3, "3")
	bst.Put(4, "4")
	bst.Put(6, "6")
	bst.Put(2, "2")

	/*
					5
				  /  \
				 3    6
			   /  \
		      2    4
	*/
	fmt.Println(bst.root.left.right.key)
	fmt.Println(bst.Get(4))

	fmt.Println(bst.Min())
	fmt.Println(bst.Max())

	fmt.Println(bst.Select(4)) // output: 排名为k的结点
	fmt.Println(bst.Rank(6))   // output: 给定key的排名

	fmt.Println("/////////////////////////////////")

	l := bst.Keys(2, 3) // 范围查找操作
	for v := l.Front(); v != nil; v = v.Next() {
		fmt.Println(v.Value)
	}

	fmt.Println("/////////////////////////////////")

	bst.Delete(5)
	fmt.Println(bst.root.key)
}

type BST struct {
	root *Node
}

type Node struct {
	key   int
	value string
	left  *Node
	right *Node
	N     int // 以该结点为根的子树中的结点总数
}

func (t *BST) Get(key int) string {
	return get(t.root, key)
}

func get(node *Node, key int) string {
	if node == nil {
		return ""
	}
	if key < node.key {
		return get(node.left, key)
	} else if key > node.key {
		return get(node.right, key)
	} else {
		return node.value
	}
}

func (t *BST) Put(key int, value string) {
	t.root = put(t.root, key, value)
}

func put(node *Node, key int, value string) *Node {
	if node == nil {
		return &Node{
			key:   key,
			value: value,
			N:     1,
		}
	}
	if key < node.key {
		node.left = put(node.left, key, value)
	} else if key > node.key {
		node.right = put(node.right, key, value)
	} else {
		node.value = value
	}
	node.N = size(node.left) + size(node.right) + 1
	return node
}

func (t *BST) Size() int {
	return size(t.root)
}

func size(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.N
	}
}

func (t *BST) Min() int {
	return min(t.root).key
}

func min(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return min(node.left)
}

func (t *BST) Max() int {
	return max(t.root).key
}

func max(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return max(node.right)
}

func (t *BST) Select(k int) int {
	return _select(t.root, k).key
}

func _select(node *Node, k int) *Node {
	if node == nil {
		return nil
	}
	t := size(node.left)
	if t > k {
		return _select(node.left, k)
	} else if t < k {
		return _select(node.right, k-t-1)
	} else {
		return node
	}
}

func (t *BST) Rank(key int) int {
	return rank(key, t.root)
}

func rank(key int, node *Node) int {
	if node == nil {
		return 0
	}
	if key < node.key {
		return rank(key, node.left)
	} else if key > node.key {
		return 1 + size(node.left) + rank(key, node.right)
	} else {
		return size(node.left)
	}
}

func (t *BST) Delete(key int) {
	t.root = _delete(t.root, key)
}

func deleteMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	node.N = size(node.left) + size(node.right) + 1
	return node
}

func _delete(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = _delete(node.left, key)
	} else if key > node.key {
		node.right = _delete(node.right, key)
	} else {
		if node.right == nil {
			return node.left
		}
		if node.left == nil {
			return node.right
		}
		t := node
		node = min(t.right)
		node.right = deleteMin(t.right)
		node.left = t.left
	}
	node.N = size(node.left) + size(node.right) + 1
	return node
}

func (t *BST) Keys(lo, hi int) *list.List {
	l := list.New()
	keys(t.root, l, lo, hi)
	return l
}

func keys(node *Node, l *list.List, lo, hi int) {
	if node == nil {
		return
	}
	if lo < node.key {
		keys(node.left, l, lo, hi)
	}
	if lo <= node.key && hi >= node.key {
		l.PushBack(node.key)
	}
	if hi > node.key {
		keys(node.right, l, lo, hi)
	}
}
