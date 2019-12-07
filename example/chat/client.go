package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"go-demo/example/chat/common"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var (
	conn   net.Conn
	active bool
)

// 客户端
func main() {
	initConn(8888)
	go receiveStdin()

	// 接收服务端消息
	for {
		var buf [512]byte
		for {
			n, _ := conn.Read(buf[0:])
			if n > 0 {
				handleServeMsg(string(buf[0:n]))
			}
		}
	}
}

// 单独的线程接收客户端输入
func receiveStdin() {

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("等待输入 ...\n")
	for input.Scan() {
		line := input.Text()
		if line == "bye" {
			break
		}
		toUser, msg := parseScannerStdin(line)

		sendMsg(toUser, msg)
		fmt.Print("等待输入 ...\n")
	}
}

// 解析命令行参数
func parseScannerStdin(line string) (to string, m string) {
	var toUser, msg string
	flag.StringVar(&toUser, "to", "", "接收方")
	flag.StringVar(&msg, "m", "", "消息内容")

	err := flag.CommandLine.Parse(strings.Split(line, " "))
	if err != nil {
		log.Print("参数解析异常！")
	}
	return toUser, msg
}

// 初始化连接
func initConn(port int) {
	addr, e := net.ResolveTCPAddr("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if e != nil {
		log.Fatal("地址解析异常！", e.Error())
	}
	conn, e = net.DialTCP("tcp", nil, addr)
	active = true
	if e != nil {
		log.Fatal("连接异常", e.Error())
	}
}

// 处理服务端消息
func handleServeMsg(msg string) {
	fmt.Print("服务端说：", msg, "\n")
}

// 发送消息
func sendMsg(toUser string, msg string) {
	if active {
		entity := common.MessageEntity{
			ToUser: toUser,
			Msg:    msg,
		}
		data, _ := json.Marshal(entity)
		_, err := conn.Write([]byte(data))
		if err != nil {
			active = false
			log.Print("消息发送异常！", err.Error())
		} else {
			log.Print("消息发送成功")
		}
	}
}
