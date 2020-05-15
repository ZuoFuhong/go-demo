package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	addr, e := net.ResolveTCPAddr("tcp", "127.0.0.1:8099")
	if e != nil {
		log.Panic(e)
	}
	listener, e := net.ListenTCP("tcp", addr)
	if e != nil {
		log.Panic(e)
	}
	for {
		conn, e := listener.AcceptTCP()
		if e != nil {
			continue
		}
		//e = conn.SetKeepAlive(true)
		if e != nil {
			continue
		}
		log.Println("新用户: ", conn.RemoteAddr())
		go doConn(conn)
	}
}

func doConn(conn *net.TCPConn) {
	for {
		err := conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		if err != nil {
			return
		}
		bytes := make([]byte, 1024)

		_, err = conn.Read(bytes)
		if err != nil {
			fmt.Println("读取超时")
			return
		}
		fmt.Println(string(bytes))
	}
}
