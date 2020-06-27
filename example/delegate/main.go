// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
)

type IntSet struct {
	data map[int]bool
	undo Undo
}

type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("no functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // Free closure for garbage collection
	}
	*undo = functions[:index]
	return nil
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}

func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

// 委托模式
// 1.先声明一个 Undo[] 的函数数组（其实是一个栈）
// 2.并实现一个通用 Add()。其需要一个函数指针，并把这个函数指针存放到 Undo[] 函数数组中。
// 3.在 Undo() 的函数中，我们会遍历Undo[]函数数组，并执行之，执行完后就弹栈。
func main() {
	ints := NewIntSet()
	for _, i := range []int{1, 3, 5, 7} {
		ints.Add(i)
	}
	fmt.Println(ints)

	for {
		if err := ints.Undo(); err != nil {
			break
		}
	}
	fmt.Println(ints)
}
