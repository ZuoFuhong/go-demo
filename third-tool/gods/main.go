package main

import (
	"fmt"

	"github.com/emirpasic/gods/maps/hashmap"

	"github.com/emirpasic/gods/lists/arraylist"
)

/*
	数据结构（Lists、Sets、Stacks、Maps、Trees），https://github.com/emirpasic/gods
*/
func main() {
	list := arraylist.New()
	list.Add(123)
	list.Add("dazuo")
	bytes, _ := list.ToJSON()
	fmt.Println(string(bytes))

	m := hashmap.New()
	m.Put("name", "dazuo")
	m.Put("age", 22)
	bytes2, _ := m.ToJSON()
	fmt.Println(string(bytes2))
}
