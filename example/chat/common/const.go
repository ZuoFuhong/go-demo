package common

// 通信的消息体
type MessageEntity struct {
	ToUser string `json:"toUser"` // 接收方
	Msg    string `json:"msg"`    // 消息体
}
