// main.go
package main

import (
	"demo2/Client"
	"demo2/EndPoint"
	"demo2/Transport"
	"fmt"
)

// 调用我们在client封装的函数就好了
func main() {
	//i, err := Client.Direct("GET", "http://127.0.0.1:8000", Transport.ByeEncodeRequestFunc, Transport.ByeDecodeResponseFunc, EndPoint.HelloRequest{Name: "songzhibin"})
	i, err := Client.ServiceDiscovery("GET", "http://127.0.0.1:8500", Transport.ByeEncodeRequestFunc, Transport.ByeDecodeResponseFunc, EndPoint.HelloRequest{Name: "songzhibin"}, "测试", true, "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, ok := i.(EndPoint.HelloResponse)
	if !ok {
		fmt.Println("no ok")
		return
	}
	fmt.Println(res)
}
