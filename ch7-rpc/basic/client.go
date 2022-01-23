package main

import (
	service "ch7-rpc/basic/string-service"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//建立http客户端
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	stringReq := &service.StringRequest{"A", "B"}
	// Synchronous call
	var reply string
	//通过call方法调用远程服务的对应方法
	err = client.Call("StringService.Concat", stringReq, &reply)
	if err != nil {
		log.Fatal("StringService error:", err)
	}
	fmt.Printf("StringService Concat : %s concat %s = %s\n", stringReq.A, stringReq.B, reply)

	stringReq = &service.StringRequest{"ACD", "BDF"}
	call := client.Go("StringService.Diff", stringReq, &reply, nil)
	_ = <-call.Done
	fmt.Printf("StringService Diff : %s diff %s = %s\n", stringReq.A, stringReq.B, reply)

}
