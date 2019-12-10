package common

import (
	"encoding/json"
	"fmt"
	"go-demo/example/chat/util"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

const (
	ReadDeadline  = 10 * time.Minute
	WriteDeadline = 10 * time.Second
)

const (
	TypeLen            = 2   // 消息类型字节数组长度
	LenLen             = 2   // 消息长度字节数据长度
	ReadContentMaxLen  = 252 // 读缓存区内容最大长度
	WriteContentMaxLen = 508 // 写缓存区内容最大长度
)

var codecFactory = NewCodecFactory(TypeLen, LenLen, ReadContentMaxLen, WriteContentMaxLen)

// 连接的上下文
type ConnContext struct {
	Codec  *Codec // 编解码器
	userId string // 用户标识
}

func NewConnContext(conn *net.TCPConn) *ConnContext {
	codec := codecFactory.GetCodec(conn)
	return &ConnContext{Codec: codec}
}

func (c *ConnContext) DoConn() {
	defer util.RecoverPanic()

	for {
		err := c.Codec.Conn.SetReadDeadline(time.Now().Add(ReadDeadline))
		if err != nil {
			c.HandleReadErr(err)
			return
		}
		_, err = c.Codec.Read()
		if err != nil {
			c.HandleReadErr(err)
			return
		}
		for {
			pack, ok, err := c.Codec.Decode()
			// 解码错误，需要中断连接
			if err != nil {
				c.Release()
				return
			}
			if ok {
				c.HandlePackage(pack)
				continue
			}
			break
		}

	}
}

// HandlePackage 处理消息包
func (c *ConnContext) HandlePackage(p *Package) {
	switch PackageType(p.Code) {
	case PackageType_PT_HEARTBEAT:
		fmt.Println("收到客户端心跳：" + string(p.Content))
		err := c.Codec.Encode(Package{Code: int(PackageType_PT_HEARTBEAT), Content: []byte("PONG")}, 10*time.Second)
		if err != nil {
			log.Println("心跳响应异常")
		}
	case PackageType_PT_MESSAGE:
		messageEntity := MessageEntity{}
		err := json.Unmarshal(p.Content, &messageEntity)
		if err != nil {
			log.Println("解析数据体异常")
		}
		c.TranspondToUser(&messageEntity)
	default:
		fmt.Print("无法处理的消息类型")
	}
}

// HandleReadErr 读取conn错误
func (c *ConnContext) HandleReadErr(err error) {
	str := err.Error()
	// 服务端主动关闭连接
	if strings.HasSuffix(str, "use of closed network connection") {
		return
	}
	c.Release()
	// 客户端主动关闭连接或者异常程序退出
	if err == io.EOF {
		return
	}
	// SetReadDeadline 之后，超时返回的错误
	if strings.HasSuffix(str, "i/o timeout") {
		return
	}
}

func (c *ConnContext) Release() {
	// 关闭tcp连接
	err := c.Codec.Release()
	if err != nil {
		log.Print(err)
	}
}

// 转发消息给指定用户
func (c *ConnContext) TranspondToUser(messageEntity *MessageEntity) {
	connContext := UserCache[messageEntity.ToUser]
	if connContext == nil {
		fmt.Println("用户", messageEntity.ToUser, " 不在线")
		return
	}
	e := connContext.Codec.Encode(Package{Code: int(PackageType_PT_MESSAGE), Content: []byte(messageEntity.Msg)}, 10*time.Second)
	if e != nil {
		fmt.Println("转发消息失败")
		return
	}
	fmt.Println("转发消息成功 toUser = ", messageEntity.ToUser)
}
