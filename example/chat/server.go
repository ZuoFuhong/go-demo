package main

import (
	"encoding/json"
	"fmt"
	"go-demo/example/chat/common"
	"log"
	"net"
)

var (
	userCache map[string]net.Conn
)

// 服务端
func main() {
	addr, e := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	checkError(e)
	listener, e := net.ListenTCP("tcp", addr)
	checkError(e)
	fmt.Print("server run in 127.0.0.1:8888\n")
	for {
		conn, e := listener.Accept()
		if e != nil {
			continue
		}
		go handleClient(conn)
	}
}

// 处理客户端消息
func handleClient(conn net.Conn) {
	userOnline(conn)

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		messageEntity := new(common.MessageEntity)
		err = json.Unmarshal(buf[0:n], &messageEntity)
		if err != nil {
			log.Print("参数解析异常！")
			return
		}
		transpondToUser(messageEntity)
	}
}

// 新用户上线
func userOnline(conn net.Conn) {
	if userCache == nil {
		userCache = make(map[string]net.Conn)
	}
	addr := conn.RemoteAddr()
	userCache[addr.String()] = conn
	log.Printf("新用户上线：%s", addr.String())
}

func checkError(err error) {
	if err != nil {
		log.Fatal("服务端启动失败！", err.Error())
		return
	}
}

// 转发消息给指定用户
func transpondToUser(messageEntity *common.MessageEntity) {
	conn := userCache[messageEntity.ToUser]
	if conn != nil {
		bytes, e := json.Marshal(messageEntity)
		if e != nil {
			log.Print("发送消息发生异常！")
			return
		}
		log.Printf("转发消息%v", messageEntity)
		_, _ = conn.Write(bytes)
	}
}
