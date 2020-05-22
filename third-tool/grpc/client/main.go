package main

import (
	"context"
	"fmt"
	"go-demo/third-tool/grpc/pb"
	"google.golang.org/grpc"
)

// RPC客户端
func main() {
	conn, e := grpc.Dial("127.0.0.1:1024", grpc.WithInsecure())
	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := pb.NewDataClient(conn)
	rp, e := client.GetUser(context.Background(), &pb.UserRq{Id: 100})
	if e != nil {
		panic(e)
	}
	fmt.Println(rp)
}
