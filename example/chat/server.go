package main

import (
	"fmt"
	"go-demo/example/chat/common"
	"log"
	"net"
)

// 服务端
func main() {
	addr, e := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	checkError(e)
	listener, e := net.ListenTCP("tcp", addr)
	checkError(e)
	fmt.Print("server run in 127.0.0.1:8888\n")
	for {
		conn, e := listener.AcceptTCP()
		if e != nil {
			continue
		}
		e = conn.SetKeepAlive(true)
		context := common.NewConnContext(conn)
		userOnline(context)
		go context.DoConn()
	}
}

// 新用户上线
func userOnline(c *common.ConnContext) {
	if common.UserCache == nil {
		common.UserCache = make(map[string]*common.ConnContext)
	}
	addr := c.Codec.Conn.RemoteAddr()
	common.UserCache[addr.String()] = c
	log.Printf("新用户上线：%s", addr.String())
}

func checkError(err error) {
	if err != nil {
		log.Fatal("服务端启动失败！", err.Error())
		return
	}
}
