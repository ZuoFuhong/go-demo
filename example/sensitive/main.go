// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import "fmt"

// 敏感词查找、验证、过滤和替换 https://github.com/importcjj/sensitive
func main() {
	tree := NewTree()
	tree.Add("world")
	tree.Add("orl")
	str := tree.replace("helloworl", '*')
	fmt.Println(str)
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	t := new(Tree)
	t.root = NewNode(0)
	return t
}

func NewNode(character rune) *Node {
	return &Node{
		Character: character,
		Child:     make(map[rune]*Node, 0),
	}
}

type Node struct {
	Character rune
	isPathEnd bool
	Child     map[rune]*Node
}

// 添加敏感词
func (t *Tree) Add(word string) {
	current := t.root
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if next, ok := current.Child[c]; ok {
			current = next
		} else {
			newNode := NewNode(c)
			current.Child[c] = newNode
			current = newNode
		}
		if i == len(runes)-1 {
			current.isPathEnd = true
		}
	}
}

// 替换文本中的敏感词，返回脱敏后的字符串
func (t *Tree) replace(text string, character rune) string {
	var (
		parent = t.root
		runes  = []rune(text)
		length = len(runes)
		left   = 0
	)
	for position := 0; position < length; position++ {
		current, found := parent.Child[runes[position]]
		if !found || (!current.isPathEnd && position == length-1) {
			parent = t.root
			position = left
			left++
			continue
		}
		if current.isPathEnd && left <= position {
			for i := left; i <= position; i++ {
				runes[i] = character
			}
		}
		parent = current
	}
	return string(runes)
}
