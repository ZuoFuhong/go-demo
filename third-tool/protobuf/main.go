package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

/*
	golang 使用 protobuf 的教程: https://www.cnblogs.com/smallleiit/p/10926794.html
    Google文档：https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

	编译命令：protoc --go_out=. *.proto
*/
func main() {
	person1 := Person{
		Id:   1,
		Name: "dazuo",
		Phones: []*Phone{
			{
				Type:   PhoneType_HOME,
				Number: "18870000000",
			},
		},
	}
	book := ContactBook{}
	book.Persons = append(book.Persons, &person1)

	// 编码
	data, e := proto.Marshal(&book)
	if e != nil {
		panic(e)
	}

	// 解码
	book2 := ContactBook{}
	e = proto.Unmarshal(data, &book2)
	if e != nil {
		panic(e)
	}
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}
}
