package main

import (
	"context"
	"fmt"
	"go-demo/third-tool/grpc/common"
	"google.golang.org/grpc"
)

// RPC客户端
func main() {
	conn, e := grpc.Dial("127.0.0.1:1024", grpc.WithInsecure())
	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := common.NewDataClient(conn)
	rp, e := client.GetUser(context.Background(), &common.UserRq{Id: 100})
	if e != nil {
		panic(e)
	}
	fmt.Println(rp)
}
