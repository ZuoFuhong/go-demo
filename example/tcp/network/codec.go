package network

import (
	"encoding/binary"
	"net"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	TypeLen       = 2                       // 消息类型字节数据长度
	BodyLen       = 2                       // 消息体字节数组长度
	HeadLen       = 4                       // 消息头部字节数组长度
	ContentMaxLen = 4092                    // 消息体最大长度
	BufferMaxLen  = ContentMaxLen + HeadLen // 缓冲buffer字节数组长度
)

type Codec struct {
	Conn   *net.TCPConn
	Buffer Buffer
}

// 协议
// --------------------------------------
// | Type(2字节) | len(2字节) | body(len) |
// --------------------------------------
func NewCodec(conn *net.TCPConn) *Codec {
	return &Codec{
		Conn:   conn,
		Buffer: newBuffer(conn, BufferMaxLen),
	}
}

func (c *Codec) Read() error {
	return c.Buffer.readFromReader()
}

func (c *Codec) Encode(p Package, duration time.Duration) error {
	bodyLen := len(p.Content)
	if bodyLen > ContentMaxLen {
		return errors.New("超出最大消息长度")
	}
	buffer := make([]byte, HeadLen+bodyLen)

	binary.BigEndian.PutUint16(buffer[0:TypeLen], uint16(p.Code))
	binary.BigEndian.PutUint16(buffer[TypeLen:TypeLen+BodyLen], uint16(bodyLen))
	copy(buffer[HeadLen:], p.Content)

	e := c.Conn.SetWriteDeadline(time.Now().Add(duration))
	if e != nil {
		return e
	}
	_, e = c.Conn.Write(buffer)
	return e
}

func (c *Codec) Decode() (*Package, bool, error) {
	typeBuf, e := c.Buffer.seek(0, TypeLen)
	if e != nil {
		return nil, false, nil
	}
	lenBuf, e := c.Buffer.seek(TypeLen, BodyLen)
	if e != nil {
		return nil, false, nil
	}
	valueType := int(binary.BigEndian.Uint16(typeBuf))
	bodyLen := int(binary.BigEndian.Uint16(lenBuf))
	if bodyLen > ContentMaxLen {
		return nil, false, errors.New("数据长度不要超过" + strconv.Itoa(ContentMaxLen))
	}
	content, e := c.Buffer.read(HeadLen, bodyLen)
	if e != nil {
		return nil, false, nil
	}
	p := Package{
		Code:    valueType,
		Content: content,
	}
	return &p, true, nil
}
