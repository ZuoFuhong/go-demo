package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// UDP客户端
func main() {
	addr, e := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	if e != nil {
		panic(e)
	}
	conn, e := net.DialUDP("udp", nil, addr)
	if e != nil {
		panic(e)
	}
	defer conn.Close()

	_, e = conn.Write([]byte("hello UDPServer"))
	if e != nil {
		panic(e)
	}

	data := make([]byte, 4)
	_, e = conn.Read(data)
	if e != nil {
		panic(e)
	}
	t := binary.BigEndian.Uint32(data)
	fmt.Println(time.Unix(int64(t), 0).String())
}
