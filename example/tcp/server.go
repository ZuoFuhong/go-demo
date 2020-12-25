package main

import (
	"go_learning_notes/example/tcp/network"
	"log"
	"net"
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
		e = conn.SetKeepAlive(true)
		if e != nil {
			continue
		}
		ctx := network.NewConnContext(conn)
		go ctx.DoConn()
	}
}
