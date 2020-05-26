package main

import (
	"context"
	"fmt"
	"go-demo/third-tool/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"runtime"
)

/*
	RPC服务端

	编译命令：protoc --go_out=plugins=grpc:. *.proto
*/

type DataService struct{}

func (s *DataService) GetUser(ctx context.Context, req *pb.UserRq) (*pb.UserRp, error) {
	fmt.Println(req)

	rp := &pb.UserRp{
		Name: "welcome!",
	}
	return rp, nil
}

// 服务器端的单向调用的拦截器
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("req: ", req, " info: ", info)
	return handler(ctx, req)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lis, e := net.Listen("tcp", "127.0.0.1:1024")
	if e != nil {
		panic(e)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor))
	pb.RegisterDataServer(s, &DataService{})
	log.Print("RPC服务已开启")

	e = s.Serve(lis)
	if e != nil {
		panic(e)
	}
}
