package network

import (
	"fmt"
	"log"
	"net"
	"time"
)

type ConnContext struct {
	Codec  *Codec
	UserId int
}

func NewConnContext(conn *net.TCPConn) *ConnContext {
	ctx := ConnContext{
		Codec:  NewCodec(conn),
		UserId: 0,
	}
	return &ctx
}

func (c *ConnContext) DoConn() {
	log.Println("新用户: ", c.Codec.Conn.RemoteAddr())
	for {
		e := c.Codec.Read()
		if e != nil {
			fmt.Println(e)
			return
		}
		for {
			p, ok, e := c.Codec.Decode()
			if e != nil {
				fmt.Println(e)
				return
			}
			if ok {
				c.HandlePackage(p)
				continue
			}
			break
		}
	}
}

// 处理包消息
func (c *ConnContext) HandlePackage(p *Package) {
	switch PackageType(p.Code) {
	case PackagetypePtHeartbeat:
		fmt.Println("客户端收到心跳：", string(p.Content))
		e := c.Codec.Encode(Package{Code: int(PackagetypePtHeartbeat), Content: []byte("PONG")}, time.Second*10)
		if e != nil {
			fmt.Println("心跳响应异常！")
		}
	case PackagetypePtMessage:
		fmt.Println("客户端收到消息：", string(p.Content))
	default:
		fmt.Println("未知的消息类型")
	}
}
