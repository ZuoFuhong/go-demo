package main

import (
	"fmt"
	"net"
	"os"
)

// TCP 客户端
func main() {
	tcpAddr, e := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	checkErr(e)
	conn, e := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(e)
	_, e = conn.Write([]byte("hello server"))
	checkErr(e)
	var buf [512]byte
	n, e := conn.Read(buf[0:])
	checkErr(e)
	fmt.Println("服务端说：", string(buf[0:n]))
	_ = conn.Close()
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
