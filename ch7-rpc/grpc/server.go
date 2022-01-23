package main

import (
	string_service "ch7-rpc/grpc/string-service"
	"ch7-rpc/pb"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//建立rpc服务端
	grpcServer := grpc.NewServer()
	stringService := new(string_service.StringService)
	pb.RegisterStringServiceServer(grpcServer, stringService) //注册服务到rpc
	grpcServer.Serve(lis)
}
