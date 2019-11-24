package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	checkError(err)
	listener, err := net.ListenTCP("tcp", addr)
	checkError(err)
	fmt.Println("server run in 127.0.0.1:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		_ = conn.Close()
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		return
	}
	fmt.Println("服务端说：", string(buf[0:n]))

	_, err = conn.Write([]byte("webcome client"))
	if err != nil {
		return
	}
}

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
