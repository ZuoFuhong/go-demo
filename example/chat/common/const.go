package common

// 通信的消息体
type MessageEntity struct {
	ToUser string `json:"toUser"` // 接收方
	Msg    string `json:"msg"`    // 消息体
}

// Package 消息包
type Package struct {
	Code    int    // 消息类型
	Content []byte // 消息体
}

type PackageType int32

const (
	PackageType_PT_UNKNOWN   PackageType = 0
	PackageType_PT_SIGN_IN   PackageType = 1
	PackageType_PT_SYNC      PackageType = 2
	PackageType_PT_HEARTBEAT PackageType = 3
	PackageType_PT_MESSAGE   PackageType = 4
)

// 聊天室-用户表示映射连接
var UserCache map[string]*ConnContext
