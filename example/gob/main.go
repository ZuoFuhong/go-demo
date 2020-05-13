package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

/*
	encoding/gob包实现了高效的序列化，特别是数据结构较复杂的，结构体、数组和切片都被支持。
*/
type Student struct {
	Name    string
	Age     uint8
	Address string
}

func main() {
	student := Student{
		Name:    "dazuo",
		Age:     24,
		Address: "Wuhan",
	}

	// 序列化
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&student)
	if err != nil {
		panic(err)
	}
	fmt.Printf("序列化后：%x\n", buffer.Bytes())

	// 反序列化
	encodeStudent := buffer.Bytes()
	decoder := gob.NewDecoder(bytes.NewReader(encodeStudent))
	var student2 Student
	err = decoder.Decode(&student2)
	if err != nil {
		panic(err)
	}
	fmt.Println("反序列化之后：", student2)
}
