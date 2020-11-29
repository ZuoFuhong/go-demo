package main

import (
	"fmt"
	"log"
	"net"
	"regexp"
)

// 获取网卡IP

func main() {
	inters, err := net.Interfaces()
	if err != nil {
		log.Panic(err)
	}
	for _, inter := range inters {
		fmt.Printf("inter.name:%s\n", inter.Name)
		addrs, err := inter.Addrs()
		if err != nil {
			log.Panic(err)
		}
		if len(addrs) == 0 {
			return
		}
		valid := regexp.MustCompile("[0-9.]+")
		fields := valid.FindAllString(addrs[0].String(), -1)
		fmt.Printf("fields:%v\n", fields[0])
	}
}
