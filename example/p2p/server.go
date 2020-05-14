package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// UDP服务器
func main() {
	// 启动：./server :8082
	addr, e := net.ResolveUDPAddr("udp", os.Args[1])
	if e != nil {
		panic(e)
	}
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	peers := make([]net.UDPAddr, 0, 2)
	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])

		peers = append(peers, *remoteAddr)
		if len(peers) == 2 {
			log.Printf("进行UDP打洞,建立 %s <--> %s 的连接\n", peers[0].String(), peers[1].String())
			_, _ = listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
			_, _ = listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
			time.Sleep(time.Second * 8)
			log.Println("中转服务器退出,仍不影响peers间通信")
			return
		}
	}
}
