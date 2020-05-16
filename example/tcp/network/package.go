package network

type PackageType int32

const (
	PackagetypePtHeartbeat PackageType = 1
	PackagetypePtMessage   PackageType = 2
)

type Package struct {
	Code    int    // 消息类型
	Content []byte // 消息体
}
