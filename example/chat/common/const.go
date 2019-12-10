package common

// Package 消息包
type Package struct {
	Code    int    // 消息类型
	Content []byte // 消息体
}

// 消息体
type MessageEntity struct {
	ToUser string `json:"toUser"` // 接收方
	Msg    string `json:"msg"`    // 消息体
}

type PackageType int32

// 消息类型
const (
	PackagetypePtHeartbeat PackageType = 1
	PackagetypePtMessage   PackageType = 2
)
