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

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lis, e := net.Listen("tcp", "127.0.0.1:1024")
	if e != nil {
		panic(e)
	}
	s := grpc.NewServer()
	pb.RegisterDataServer(s, &DataService{})
	log.Print("RPC服务已开启")

	e = s.Serve(lis)
	if e != nil {
		panic(e)
	}
}
