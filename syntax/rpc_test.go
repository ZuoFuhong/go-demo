package syntax

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"testing"
)

/*
	RPC是远程过程调用的简称，是分布式系统中不同节点间流行的通信方式。在互联网时代，RPC已经和IPC一样成为一个不可或缺的基础构件。
	因此Go语言的标准库也提供了一个简单的RPC实现。

	跨语言RPC
	1.标准库的RPC默认采用Go语言特有的gob编码，因此从其它语言调用Go语言实现的RPC服务将比较困难。
	2.Go语言的RPC框架有两个比较有特色的设计：一个是RPC数据打包时可以通过插件实现自定义的编码和解码；另一个是RPC建立在抽象的
      io.ReadWriteCloser接口之上的，我们可以将RPC架设在不同的通讯协议之上。
*/

type UserSerivce struct{}

// Go语言的RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型。
func (u *UserSerivce) GetUserById(uid int, reply *string) error {
	*reply = "dazuo_" + strconv.Itoa(uid)
	return nil
}

func Test_Server(t *testing.T) {
	err := rpc.RegisterName("UserService", new(UserSerivce))
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	conn, err := listen.Accept()
	if err != nil {
		panic(err)
	}
	rpc.ServeConn(conn)
}

// 通过rpc.Dial拨号RPC服务，然后通过client.Call调用具体的RPC方法。在调用client.Call时，第一个参数是用点号链接的RPC服务名字和方法名字，
// 第二和第三个参数分别我们定义RPC方法的两个参数。
func Test_Client(t *testing.T) {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	var reply string
	err = client.Call("UserService.GetUserById", 1, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
