// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Map(t *testing.T) {
	// 使用make创建map
	s := make(map[string]string)
	s["name"] = "dazuo"
	s["age"] = "22"
	fmt.Printf("%#v, name = %s, len = %d\n", s, s["name"], len(s))

	// 判断key是否存在map中
	val, prs := s["name"]
	fmt.Printf("%s, %t\n", val, prs) // output: dazuo, true

	// 声明map时填充元素
	m := map[string]string{
		"W": "forward",
		"A": "left",
	}
	// 遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 删除map中的键值对
	// 清空map的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
	delete(m, "W")
	fmt.Println(m)
}

func Test_syncMap(t *testing.T) {
	// 并发环境中的map
	var scene sync.Map

	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

// 将map作为函数的参数进行传递
func Test_map_param(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "dazuo"

	fmt.Println(m)
	changeMap(m)
	fmt.Println(m)
}

func changeMap(param map[string]string) {
	param["name"] = "locy"
}
