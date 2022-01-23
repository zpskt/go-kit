package main

import (
	"ch7-rpc/basic/string-service"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	//生成StringService结构体
	stringService := new(service.StringService)
	//使用rpc.Register注册服务
	registerError := rpc.Register(stringService)
	if registerError != nil {
		log.Fatal("Register error: ", registerError)
	}
	rpc.HandleHTTP()
	//用net.Listen监听对应socket并对外提供服务
	l, e := net.Listen("tcp", "127.0.0.1:1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
