package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"go-demo/example/chat/common"
	"go-demo/example/chat/util"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	codec       *common.Codec
	codeFactory = common.NewCodecFactory(2, 2, 65536, 1024)
)

// 客户端
func main() {
	conn := initConn(8888)

	codec = codeFactory.GetCodec(conn)
	go receiveStdin()
	go Heartbeat()

	// 接收服务端消息
	for {
		_, e := codec.Read()
		if e != nil {
			fmt.Println(e)
			return
		}

		for {
			pack, ok, e := codec.Decode()
			if e != nil {
				fmt.Print(e)
				return
			}
			if ok {
				handlePackage(*pack)
				continue
			}
			break
		}
	}
}

// 单独的线程接收客户端输入
func receiveStdin() {
	input := bufio.NewScanner(os.Stdin)

	var toUser, msg string
	flag.StringVar(&toUser, "to", "", "接收方")
	flag.StringVar(&msg, "m", "", "消息内容")

	fmt.Println("开始接受控制台输入...")
	for input.Scan() {
		line := input.Text()
		if line == "bye" {
			break
		}
		parseScannerStdin(line)

		if toUser == "" {
			log.Println("接受人不能为空")
			continue
		}
		if msg == "" {
			log.Println("不能发空消息")
			continue
		}
		sendMsg(toUser, msg)
		fmt.Println("等待输入 ...")
	}
}

// 解析命令行参数
func parseScannerStdin(line string) {
	defer util.RecoverPanic()
	err := flag.CommandLine.Parse(strings.Split(line, " "))
	if err != nil {
		log.Print("参数解析异常！")
	}
}

// 初始化连接
func initConn(port int) net.Conn {
	addr, e := net.ResolveTCPAddr("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if e != nil {
		log.Fatal("地址解析异常！", e.Error())
	}
	conn, e := net.DialTCP("tcp", nil, addr)
	if e != nil {
		log.Fatal("连接异常", e.Error())
	}
	return conn
}

// 处理服务端消息
func handlePackage(p common.Package) {
	switch common.PackageType(p.Code) {
	case common.PackageType_PT_HEARTBEAT:
		fmt.Println("收到服务端心跳响应：" + string(p.Content))
	case common.PackageType_PT_MESSAGE:
		fmt.Println("客户端收到消息：" + string(p.Content))
	default:
		fmt.Println("无法处理的消息类型")
	}
}

// 发送消息
func sendMsg(toUser string, msg string) {
	entity := common.MessageEntity{
		ToUser: toUser,
		Msg:    msg,
	}
	data, _ := json.Marshal(entity)
	err := codec.Encode(common.Package{Code: int(common.PackageType_PT_MESSAGE), Content: []byte(data)}, 10*time.Second)
	if err != nil {
		log.Print("消息发送异常！", err.Error())
	} else {
		log.Print("消息发送成功")
	}
}

// 发送心跳
func Heartbeat() {
	ticker := time.NewTicker(time.Second * 15)
	for range ticker.C {
		err := codec.Encode(common.Package{Code: int(common.PackageType_PT_HEARTBEAT), Content: []byte("PING")}, 10*time.Second)
		if err != nil {
			print(err)
		}
	}
}
