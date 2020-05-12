package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// UDP服务端
func main() {
	addr, e := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	if e != nil {
		panic(e)
	}
	conn, e := net.ListenUDP("udp", addr)
	if e != nil {
		panic(e)
	}
	defer conn.Close()

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	_, remoteAddr, e := conn.ReadFromUDP(data)
	if e != nil {
		panic(e)
	}
	daytime := time.Now().Unix()
	fmt.Println(remoteAddr, string(data))

	reply := make([]byte, 4)
	binary.BigEndian.PutUint32(reply, uint32(daytime))
	_, _ = conn.WriteToUDP(reply, remoteAddr)
}
