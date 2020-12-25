package main

import (
	"context"
	"fmt"
	pb "go_learning_notes/third-tool/grpc/proto"
	"testing"

	"google.golang.org/grpc"
)

func Test_Client(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:1024", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewBasicDataSvrClient(conn)
	rpcRsp, err := client.GetUserInfo(context.Background(), &pb.GetUserInfoReq{Id: 100})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("rpcRsp = %v\n", rpcRsp)
}
