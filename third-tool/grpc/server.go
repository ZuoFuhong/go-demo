package main

import (
	"context"
	pb "go_learning_notes/third-tool/grpc/proto"
	"log"
	"net"
	"runtime"

	"google.golang.org/grpc"
)

type BasicDataInfoService struct{}

func (s *BasicDataInfoService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoRsp, error) {
	log.Printf("reqData{id:%d}", req.GetId())
	rsp := &pb.GetUserInfoRsp{
		Name: "mars",
		Age:  24,
		City: "Wuhan",
	}
	return rsp, nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lis, err := net.Listen("tcp", "127.0.0.1:1024")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterBasicDataSvrServer(s, &BasicDataInfoService{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
