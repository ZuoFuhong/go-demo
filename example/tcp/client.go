package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	raddr, e := net.ResolveTCPAddr("tcp", "127.0.0.1:8099")
	if e != nil {
		log.Panic(e)
	}

	conn, e := net.DialTCP("tcp", nil, raddr)
	if e != nil {
		panic(e)
	}
	fmt.Println(conn.RemoteAddr().String()) // output: 127.0.0.1:8099
	for i := 0; i < 5; i++ {
		_, e = conn.Write([]byte("hello server"))
		if e != nil {
			panic(e)
		}
		time.Sleep(time.Second * 2)
	}

	time.Sleep(time.Hour)
}
