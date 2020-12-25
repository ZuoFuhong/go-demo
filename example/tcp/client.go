package main

import (
	"fmt"
	"go_learning_notes/example/tcp/network"
	"log"
	"net"
	"time"
)

var (
	codec *network.Codec
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
	go heartbeat()

	codec = network.NewCodec(conn)
	for {
		e := codec.Read()
		if e != nil {
			fmt.Println(e)
			return
		}
		for {
			p, b, e := codec.Decode()
			if e != nil {
				fmt.Println(e)
				return
			}
			if b {
				handlePackage(p)
				continue
			}
			break
		}
	}
}

// 处理包消息
func handlePackage(p *network.Package) {
	switch network.PackageType(p.Code) {
	case network.PackagetypePtHeartbeat:
		fmt.Println("客户端收到心跳：", string(p.Content))
	case network.PackagetypePtMessage:
		fmt.Println("客户端收到消息：", string(p.Content))
	default:
		fmt.Println("未知的消息类型")
	}
}

// 发送心跳
func heartbeat() {
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		p := network.Package{Code: int(network.PackagetypePtHeartbeat), Content: []byte("PING")}
		e := codec.Encode(p, 10*time.Second)
		if e != nil {
			log.Panic(e)
		}
	}
}
